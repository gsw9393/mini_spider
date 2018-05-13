package data_struct

import (
    "parse"
)

type UrlTask struct {
    url    string
    depth  string
}

func NewUrlTask(url string, depth int) *UrlTask {
    urltask := new(UrlTask)
    urltask.url = url
    urltask.depth = depth

    return urltask
}

type MainSpider struct {
    synqueue *data_struct.SynQueue
    urlmap   *data_struct.UrlMap
    confinfo *parse.ConfInfo
    workers    []*data_struct.Worker
}

func NewMainSpider(confinfo *parse.ConfInfo, urls []string) (*MainSpider) {
    mainspider := new(MainSpider)

    mainspider.confinfo = confinfo
    mainspider.synqueue = NewSynQueue()
    mainspider.urlmap = NewUrlMap()

    for _, urlLink := range urls {
        task := NewUrlTask(urlLink, 0)
        mainspider.synqueue.Add(task)
    }

    for i := 0; i < confinfo.ThreadCount; i++ {
        worker := NewWorker(mainspider.synqueue, mainspider.urlmap, confinfo, confinfo.TargetUrl)
        mainspider.workers = append(mainspider.workers, workers)
    }

    return mainspider
}

func (mainspider *MainSpider) Start() {
    for _, worker := range mainspider.workers {
        go worker.Start()
    }
}

