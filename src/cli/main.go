package main

import (
    "fmt"
    "time"
    "os"
    "model"
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
    p, pause: pause playing
    s, skip: skip the song playing now
    r, rec: show recommand channel
    ls, list: show all channel
    c %, channel %: change channel, type id or name both ok. example:
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
    manager := model.GetManagerInstance()

    PrintHelp()
    fmt.Print(">> ")

    go func() {
        fmt.Println("\rAlways Online - 林俊杰 00:11\b2")
        time.Sleep(1e9 * 10)
    }()

    go func() {
        for {
            cmd := <-ch
            // Split cmd into words, then cmd is the first word
            switch cmd {
                case "h":
                    fallthrough
                case "help":
                    PrintHelp()
                case "q":
                    fallthrough
                case "quit":
                    os.Exit(0)
                case "f":
                    fallthrough
                case "fav":
                    manager.Player().Fav()
                case "u":
                    fallthrough
                case "unfav":
                    manager.Player().UnFav()
                case "d":
                    fallthrough
                case "del":
                    manager.Player().Del()
                case "l":
                    fallthrough
                case "loop":
                    manager.Player().Loop()
                case "p":
                    fallthrough
                case "pause":
                    manager.Player().Pause()
                case "s":
                    fallthrough
                case "skip":
                    manager.Player().Skip()
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
