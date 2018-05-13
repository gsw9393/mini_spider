package global_data

// 配置文件数据的结构体
type confInfo struct {
    //配置文件通过tag指定配置文件中的名称
    UrlListFile     string `ini:"urlListFile"`
    OutputDirectory string `ini:"outputDirectory"`
    MaxDepth        int    `ini:"maxDepth"`
    CrawInterval    int    `ini:"crawInterval"`
    CrawTimeout     int    `ini:"crawTimeout"`
    TargetUrl       string `ini:"targetUrl"`
    ThreadCount     int    `ini:"threadCount"`
}


