package main

import (
    "os"
    "fmt"
    "flag"

    "parse"
    "util"
    log "github.com/log4go"
)

var (
    help      bool
    version   bool
    confPath  string
    logPath   string
)

func init() {

    flag.BoolVar(&help,       "h", false,       "this help")
    flag.BoolVar(&version,    "v", false,       "show version and exit")
    flag.StringVar(&confPath, "c", "../conf/spider.conf",   "show configuration file path `confpath`")
    flag.StringVar(&logPath,  "l", "../log",    "show log file path `logpath`")
    flag.Usage = usage  //改变默认的Usage
}

func usage() {
    fmt.Fprintf(os.Stderr, `mini_spider version: mini_spider/1.0.0
Usage: ./mini_spider [-h help] [-v version] [-c confPath] [-l logPath]
    
options:
`)
    flag.PrintDefaults()
}


func main() {
    flag.Parse()

    if help {
        flag.Usage()
        return
    }

    err := util.MakirPath(logPath)
    if err != nil {
        fmt.Println("mkdir error")
        return
    }
    //应该设置为全局变量
    logFileName := "mini_spider.conf"
    /* set log */
    log.AddFilter("file", log.DEBUG, log.NewFileLogWriter(logPath + "/" + logFileName, false))
    defer log.Close()
    //log.Log(log.ERROR, "main", "hello log")

    /* configuration read */
    config, err := parse.ReadConfig(confPath)
    if err != nil {
        fmt.Println("configuration file read error")
        return
    }
    /*
    fmt.Println(config.UrlListFile)
    fmt.Println(config.OutputDirectory)
    fmt.Println(config.MaxDepth)
    fmt.Println(config.CrawInterval)
    fmt.Println(config.CrawTimeout)
    fmt.Println(config.TargetUrl)
    fmt.Println(config.ThreadCount)
    */

    seeds, err := parse.DecodeSeed(config.UrlListFile)
    if err != nil {
        fmt.Println("decode seed error")
        return
    }
    /*
    for _, seed := range seeds {
        fmt.Println(seed)
    }
    */



}
