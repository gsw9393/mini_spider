package util

import (
    "fmt"
    "path"
    "io/ioutil"
    "net/url"
    "net/http"
    "os"
    log "github.com/log4go"
)

func MakirPath(path string) error {
    _, err := os.Stat(path)
    if err == nil {
        return nil
    }
    if os.IsNotExist(err) {
        err = os.MkdirAll(path, os.ModePerm)
        if err != nil {
            fmt.Println("mdkir dir error")
            return err
        }
    }
    return nil
}

func Save2Disk(urllink string, outputPath string, resp *http.Response) error {

    filepath := path.Join(outputPath, url.QueryEscape(urllink))

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Error("ioutil.ReadAll error with urllink: %s", urllink)
        return err
    }

    err = ioutil.WriteFile(filepath, []byte(body), 0644)
    if err != nil {
        log.Error("write %s error", filepath)
        return err
    }
    return nil
}
