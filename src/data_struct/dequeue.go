package data_struct

import (
    "sync"
    "github.com/queue"
)

type SynQueue struct {
    lock     sync.Mutex
    popable  *sync.Cond
    buffer   *queue.Queue
    closed   bool
}

func NewSynQueue() *SynQueue {
    ch := &SynQueue{
        buffer: queue.New(),
    }
    ch.popable = sync.NewCond(&ch.lock)
    return ch
}

func (ch *SynQueue) Pop() (v interface{}) {
    flag := ch.popable
    buffer := ch.buffer

    ch.lock.Lock()
    defer ch.lock.Unlock()

    for buffer.Length() == 0 && !ch.closed {
        flag.Wait()
    }

    if buffer.Length() > 0 {
        v = buffer.Peek()
        buffer.Remove()
    }

    return
}

func (ch *SynQueue) TryPop() (v interface{}, ok bool) {
    buffer := ch.buffer

    ch.lock.Lock()
    defer ch.lock.Unlock()

    if buffer.Length() > 0 {
        v = buffer.Peek()
        buffer.Remove()
        ok = true
    } else if ch.closed {
        ok = true
    }
    return
}

func (ch *SynQueue) Push(v interface{}) {
    ch.lock.Lock()
    defer ch.lock.Unlock()

    if !ch.closed {
        ch.buffer.Add(v)
        ch.popable.Signal()
    }
}

func (ch *SynQueue) Len() (l int) {
    ch.lock.Lock()
    defer ch.lock.Unlock()

    l = ch.buffer.Length()
    return
}

func (ch *SynQueue) Close() {
    ch.lock.Lock()
    defer ch.lock.Unlock()

    if !ch.closed {
        ch.closed = true
        ch.popable.Signal()
    }
}
