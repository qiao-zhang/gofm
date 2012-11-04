package main

import (
    "fmt"
    "time"
    "os"
)

func PrintHelp(){
    fmt.Println(`Gofm - A Douban.fm Command Line Interface

Github: 
    http://dou.bz/3c8G7L
    等你的pull request ;)

Useage:
    h, help: list help
    q, quit: quit Gofm
    f, fav: fav the song playing now
    u, unfav: unfav the song playing now
    d, del: move the song playing now into trash
    l, loop: loop playing current song
    s, skip: skip the song playing now
    r, rec: show recommand channel
    ls, list: show all channel
    c %, choose %: change channel, type id or name both ok. example:
        >> l
        华语(1) 欧美(2)
        >> c 1
        >> c 华语`)
}

func PrintNotSupport() {
    fmt.Println(`not supported cmd ;(`)
}

func main() {
    ch := make(chan string)
    PrintHelp()
    fmt.Print(">> ")

    go func() {
        fmt.Println("\rAlways Online - 林俊杰 00:11\b2")
        time.Sleep(1e9 * 10)
    }()

    go func() {
        for {
            cmd := <-ch
            switch cmd {
                case "h":
                    fallthrough
                case "help":
                    PrintHelp()
                case "q":
                    fallthrough
                case "quit":
                    os.Exit(0)
                case "":
                    // do nothing
                default:
                    PrintNotSupport()
            }
            fmt.Print(">> ")
        }
    }()

    go func() {
        var cmd string
        for {
            fmt.Scanln(&cmd)
            ch <- cmd
            cmd = ""
        }
    }()

    for {
        time.Sleep( time.Second )
    }
}
