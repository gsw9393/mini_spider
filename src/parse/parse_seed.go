package parse

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)



func DecodeSeed(seedPath string) ([]string, error) {
    var seeds []string
    data, err := ioutil.ReadFile(seedPath)
    if err != nil {
        fmt.Println("read seed file error")
        return seeds, err
    }
    err = json.Unmarshal(data, &seeds)
    if err != nil {
        fmt.Println("json Unmarshal error")
        return seeds, err
    }
    return seeds, nil
}

