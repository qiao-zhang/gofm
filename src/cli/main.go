package main

import (
    "fmt"
    "model"
)


func main() {
    c := &model.Creator{"doubanfm", "http://douban.fm"}
    fmt.Println( c )

    ch := model.NewChannel("Huayu", "doubanfm", 1024, c)
    fmt.Println( ch )
}
