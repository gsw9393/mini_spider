package data_struct

import (
    "sync"
)

type UrlMap struct {
    urlmap map[string]bool
    mutex  sync.Mutex
}

func NewUrlMap() *UrlMap {
    um := new(UrlMap)
    um.urlmap = make(map[string]bool)
    um.mutex = sync.Mutex{}
    return um
}

func (um *UrlMap) Add(url string) {
    um.mutex.Lock()
    defer um.mutex.Unlock()
    um.urlmap[url] = true
}

func (um *UrlMap) IsExist(url string) bool {
    um.mutex.Lock()
    defer um.mutex.Unlock()
    _, ok := um.urlmap[url]
    return ok
}
