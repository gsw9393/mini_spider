package parse

import (
    "fmt"
    "gopkg.in/ini.v1"
)

//配置文件通过tag指定配置文件中的名称
type ConfInfo struct {
    UrlListFile     string `ini:"urlListFile"`
    OutputDirectory string `ini:"outputDirectory"`
    MaxDepth        int    `ini:"maxDepth"`
    CrawInterval    int    `ini:"crawInterval"`
    CrawTimeout     int    `ini:"crawTimeout"`
    TargetUrl       string `ini:"targetUrl"`
    ThreadCount     int    `ini:"threadCount"`
}

func confCheck(config ConfInfo) error {

    if config.UrlListFile == "" {
        return  fmt.Errorf("urlListFile is empty")
    }

    if config.OutputDirectory == "" {
        return  fmt.Errorf("outputDirectory is empty")
    }

    if config.MaxDepth <= 0 {
        return  fmt.Errorf("maxDepth's value is wrong")
    }

    if config.CrawInterval <= 0 {
        return  fmt.Errorf("crawInterval's value is wrong")
    }

    if config.CrawTimeout <= 0 {
        return  fmt.Errorf("crawTimeout's value is wrong")
    }

    if config.TargetUrl == "" {
        return  fmt.Errorf("targetUrl's value is wrong")
    }

    if config.ThreadCount <= 0 {
        return  fmt.Errorf("threadCount's value is wrong")
    }

    return nil
}

//将配置文件读取到结构体中
func ReadConfig(path string) (ConfInfo, error) {

    var config ConfInfo
    //加载配置文件
    conf, err := ini.Load(path)
    if err != nil {
        fmt.Println("load config error")
        return config, err
    }
    conf.BlockMode = false
    //解析成结构体
    err = conf.MapTo(&config)
    if err != nil {
        fmt.Println("mapto config error")
        return config, err
    }

    err = confCheck(config)
    if err != nil {
        fmt.Println("configuration file's parameters is wrong")
        return config, err
    }
    return config, nil
}

/*
func main() {

    config, err := ReadConfig("../../conf/spider.conf")
    if err != nil {
        fmt.Println("error")
        return
    }

    fmt.Println(config.UrlListFile)
    fmt.Println(config.OutputDirectory)
    fmt.Println(config.MaxDepth)
    fmt.Println(config.CrawInterval)
    fmt.Println(config.CrawTimeout)
    fmt.Println(config.TargetUrl)
    fmt.Println(config.ThreadCount)
}
*/
