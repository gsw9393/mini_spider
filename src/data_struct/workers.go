package data_struct

import (
    "regexp"
    "parse"
    "util"
    log "github.com/log4go"
)

type Worker struct {
    synqueue *SynQueue
    urlmap   *UrlMap
    confinfo *parse.ConfInfo
    pattern  *regexp.Regexp
}

func NewWorker(synqueue *SynQueue, urlmap *UrlMap, confinfo *ConfInfo, url string) *Worker {
    worker := new(Worker)
    worker.synqueue = synqueue
    worker.urlmap   = urlmap
    worker.confinfo = confinfo
    work.pattern, _ = regexp.Compile(url)

    return work
}

func (worker *Worker) Start() {

    urlTask := worker.synqueue.Pop()

    client := parse.GetClient(timeout)
    resp, err := client.Get(urlTask.url)
    defer resp.Body.Close()
    if err != nil {
        log.Error("http Get error with urllink: %s", urlTask.url)
        continue
    }

    if work.pattern.MatchString(urlTask.url) {
        err = util.Save2disk(urlTask.url, worker.confinfo.OutputDirectory, resp)
        if err != nil {
            log.Error("Save2disk error with urlLink: %s", urlTask.url)
        }
    }

    worker.urlmap.Add(urlTask.url)

    if urlTask.depth < worker.confinfo.MaxDepth {
        urlLinks, err := GetUrlLinks(resp, worker.confinfo.CrawTimeout, work.pattern)
        if err != nil {
            log.Error("GetUrlLinks error with url link: %s", urlTask.url )
        }

        for _,  href := range urlLinks {
            newtask := UrlTask{url: href, depth: urlTask.depth + 1}
            worker.synqueue.Push(newtask)
        }
    }
}
