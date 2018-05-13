package parse

import (
    "io"
    //"fmt"
    "net"
    "time"
    "regexp"
    "net/http"
    log "github.com/log4go"
    "github.com/PuerkitoBio/goquery"
    "golang.org/x/net/html/charset"
    "golang.org/x/text/encoding/htmlindex"
)

func detectContentCharset(body []byte) string {
    if _, name, ok := charset.DetermineEncoding(body, ""); ok {
        return name
    }
    return "utf-8"
}

func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
    if (charset == "") {
        buf := make([]byte, 1024)
        body.Read(buf)
        charset = detectContentCharset(buf)
    }
    e, err := htmlindex.Get(charset)
    if err != nil {
        return nil, err
    }
    if name, _ := htmlindex.Name(e); name != "utf-8" {
        body = e.NewDecoder().Reader(body)
    }
    return body, nil
}

func GetClient(timeout int) *http.Client {
    client := http.Client{
        Transport: &http.Transport{
            Dial: func(netw, addr string) (net.Conn, error) {
                client, err := net.DialTimeout(netw, addr, time.Second * time.Duration(timeout))
                if err != nil {
                    return nil, err
                }
                client.SetDeadline(time.Now().Add(1 * time.Second))
                return client, nil
            },
        },
    }
    return &client
}



/*
func GetHtmlBody(timeout int, url string) (, error) {

    client := GetClient(timeout)

    resp, err := client.Get(url)
    if err != nil {
        log.Error("GetHtmlBody error, url:%s", url)
        return  nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        err := fmt.Errorf("client Get error with StatusCode: %d", resp.StatusCode)
        return nil, err
    }

    reader, err := DecodeHTMLBody(resp.Body, "")
    if err != nil {
        fmt.Println("parse.DecodeHTMLBody error")
        log.Log(log.ERROR, "DecodeHTMLBody", "parse DecodeHTMLBody error")
    }
    return &reader, nil
}

*/
func GetUrlLinks(resp *http.Response, timeout int, regexp *regexp.Regexp) ([]string, error) {

    doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil{
        log.Error("goquery NewDocumentFromReader error")
        return nil, err
    }
    var urlLinks = make([]string, 0)
    doc.Find("a").Each(func(_ int, link *goquery.Selection) {
        href, ok := link.Attr("href")
        if ok  {
            if regexp.MatchString(href){
                urlLinks = append(urlLinks, href)
            }
        }
    })
    return urlLinks, nil

