package model

import (
    "time"
    "net/http"
    "log"
    "strconv"
    "encoding/json"
    "io/ioutil"
)

type Playlist struct {
    R int
    Err string
    Song []Song
}

func NewPlaylist(r int, msg string, song []Song) *Playlist {
    return &Playlist{r, msg, song}
}

func (this *Playlist) FetchChannel(channel int, typ string) {
    this.FetchChannelBase(channel, typ, "")
}

func (this *Playlist) FetchChannelNextSong( channel int, typ string, sid string) {
    this.FetchChannelBase(channel, typ, sid)
}

func (this *Playlist) FetchChannelBase(channel int, typ string, sid string) {
    fetch_url := fm_site + "/j/mine/playlist?type=" + typ + "&channel=" + strconv.Itoa(channel)
    if sid != "" {
        fetch_url += "&sid=" + sid
    }
    chs := make(chan int)
    timeout := make(chan bool, 1)
    if this == nil {
        this = new(Playlist)
    }

    go func() {
        time.Sleep( time.Second * 10 )
        timeout <-true
    }()

    go func() {
        resp, err := http.Get(fetch_url)
        if err != nil {
            log.Fatal("Gofm can't fetch playlist, plz check network")
            return
        }
        defer resp.Body.Close()

        playlist_str, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal("Gofm can't read from playlist, plz contact developer")
            return
        }

        parse_err := json.Unmarshal(playlist_str, this)
        if parse_err != nil {
            panic( parse_err)
            //log.Fatal("Gofm can't parse playlist, plz contact developer")
        }

        if this.R == fm_mine_playlist_err_handle {
            log.Fatalf("Douban.fm Error Msg: %s", this.Err)
        }

        chs <- this.R
    }()

    select {
    case <-chs:
        // wrong is nothing
    case <-timeout:
        log.Printf("Gofm fetch playlist timeout")
    }
}
