package model

import (
    "net/http"
    "time"
    "log"
    "encoding/json"
    "io/ioutil"
)


type Channel struct {
    Id int
    Name string
    Intro string
    Song_num int
    Hot_songs []string
    Creator *Creator
}

func NewChannel(id int, name, intro string, song_num int,
            hot_songs []string, creator *Creator) *Channel {
    return &Channel{id, name, intro, song_num, hot_songs, creator}
}

type HotChannels struct{
    Status bool
    Data struct{
        Total int
        Channels []Channel
    }
    Error string
}

type ChannelInfo struct {
    Status bool
    Data struct {
        Channels []Channel
    }
    Error string
}

func (c Channel) FetchHotChannels() (total int, channels []Channel) {
    fetch_url := fm_site + "/j/explore/hot_channels"

    chs := make(chan bool)
    timeout := make(chan bool, 1)

    go func() {
        time.Sleep( time.Second * 10 )
        timeout <- true
    }()

    go func() {
        resp, err := http.Get(fetch_url)
        if err != nil {
            log.Fatal("Gofm can't fetch hot channels, plz check network")
            return
        }
        defer resp.Body.Close()

        channel_str, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal("Gofm can't read from hot channels, plz contact developer")
            return
        }

        hot_channels := new(HotChannels)
        parse_err := json.Unmarshal(channel_str, hot_channels)
        if parse_err != nil {
            panic( parse_err)
        }

        total = hot_channels.Data.Total
        channels = hot_channels.Data.Channels
        chs <- hot_channels.Status
    }()

    select {
        case <- chs:
            return
        case <-timeout:
            log.Printf("Gofm fetch playlist timeout")
            return
    }
    return
}


func (c Channel) FetchChannelInfo(id string) (channel *Channel){
    defer func() {
        if r := recover(); r != nil {
            log.Print("Fatal! not found this channel.\n>> ")
            return
        }
    }()
    fetch_url := fm_site + "/j/explore/channel_info?id=" + id

    chs := make(chan bool)
    timeout := make(chan bool, 1)

    go func() {
        time.Sleep( time.Second * 10 )
        timeout <- true
    }()

    go func() {
        resp, err := http.Get(fetch_url)
        if err != nil {
            log.Fatal("Gofm can't fetch hot channels, plz check network")
            return
        }
        defer resp.Body.Close()

        channel_str, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal("Gofm can't read from hot channels, plz contact developer")
            return
        }

        channel_info := new(ChannelInfo)
        parse_err := json.Unmarshal(channel_str, channel_info)
        if parse_err != nil {
            panic( parse_err)
        }

        if len(channel_info.Data.Channels) > 0 {
            channel = &channel_info.Data.Channels[0]
        }
        chs <- channel_info.Status
    }()

    select {
        case <- chs:
            return
        case <-timeout:
            log.Printf("Gofm fetch playlist timeout")
            return
    }
    return
}


