package main

import (
    "fmt"
    "time"
    "os"
    "model"
    "strings"
    "strconv"
)

func PrintHelp(){
    fmt.Println(`Gofm - A Douban.fm Command Line Interface

Github: 
    http://dou.bz/3c8G7L
    等你的pull request ;)

Usage:
    h, help: list help
    q, quit: quit Gofm
    f, fav: fav the song playing now [ login required ]
    u, unfav: unfav the song playing now [ login required ]
    d, del: move the song playing now into trash [ login required ]
    l, loop: loop playing current song, :) I like it.
    p, pause: pause playing
    r, resume: resume playing
    s, skip: skip the song playing now
    ls, list: show all channel
    hc, hot_channels: show hot channels
    ci %, channel_info %: show channel information
    c %, channel %: change channel, type "c id" and then enter. example:
        >> ls
        华语(1) 欧美(2)
        >> c 1`)
}

func PrintNotSupport() {
    fmt.Println(`not supported cmd ;(`)
}

func PrintHotChannels() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Print("Fatal! not found hot channels.\n>> ")
            return
        }
    }()
    fmt.Print("\rFetching Hot Channels...")
    channels := new(model.Channel)
    total, hot_channels := channels.FetchHotChannels()
    if total != 0 {
        fmt.Print("\r")
    }
    for _, c := range(hot_channels) {
        fmt.Print(c.Name, "(", c.Id, ") ")
    }
    fmt.Print("\n>> ")
}

func PrintChannelInfo(id string) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Print("Fatal! not found this channel.\n>> ")
            return
        }
    }()
    fmt.Println("\rFetching Channels ", id, " ...")
    channel := new(model.Channel)
    ch := channel.FetchChannelInfo(id)
    fmt.Print(ch.Name + "(" + id + ") \n" +
            "Intro: " + ch.Intro + "\n" +
            "Hot Songs: " + strings.Join(ch.Hot_songs, " ") +
            " (totals: " + strconv.Itoa(ch.Song_num) + ")")
    fmt.Print("\n>> ")
}

func main() {
    ch := make(chan string)
    manager := model.GetManagerInstance()
    manager_trigger := make(chan string)

    PrintHelp()
    fmt.Print(">> ")


    // This goroutine is used to recieve message from manager
    go manager.Start(manager_trigger)

    go func() {
        for {
            select {
            case msg := <-manager_trigger:
                switch msg{
                case "current_song":
                    song := manager.Playlist().Song[manager.ProgressInPlaylist()]
                    fmt.Println("\r"+ song.Format() )
                case "pause_song":
                    fmt.Println("\rpausing, U can type `r` or `resume` to continue.")
                case "loop_song":
                    song := manager.Playlist().Song[manager.ProgressInPlaylist()]
                    fmt.Println( "\rlooping: " + song.Format() )
                }
            }
            fmt.Print(">> ")
        }
    }()

    // handle user input command
    go func() {
        for {
            cmd := <-ch
            cmds := strings.Split(cmd, " ")
            if len(cmds) > 0 {
                cmd = cmds[0]
            }
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
                    manager.Player().Loop(manager_trigger)
                case "p":
                    fallthrough
                case "pause":
                    manager.Player().Pause(manager_trigger)
                case "r":
                    fallthrough
                case "resume":
                    manager.Player().Resume(manager_trigger)
                case "s":
                    fallthrough
                case "skip":
                    manager.Player().Skip(manager_trigger)
                case "hc":
                    fallthrough
                case "hot_channels":
                    go PrintHotChannels()
                case "ci":
                    fallthrough
                case "channel_info":
                    if len(cmds) >= 2 {
                        go PrintChannelInfo(cmds[1])
                    } else {
                        PrintNotSupport()
                    }
                case "c":
                    fallthrough
                case "channel":
                    if len(cmds) >= 2 {
                        manager.ChooseChannel(cmds[1])
                    } else {
                        PrintNotSupport()
                    }
                case "":
                    // do nothing
                default:
                    PrintNotSupport()
            }
            fmt.Print(">> ")
        }
    }()

    // This goroutine is used to waiting user's command
    go func() {
        var cmd string
        var param string
        for {
            fmt.Scanln(&cmd, &param)
            cmd = cmd + " " + param
            ch <- cmd
            cmd = ""
        }
    }()

    for {
        time.Sleep( time.Second )
    }
}
