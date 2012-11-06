package model

import (
    "time"
    "log"
    //"fmt"
)

// Mp3Player 的关注焦点在正确的播放传入的mp3文件
// 它与Manager需要的交互是发送当前的播放状态
type Mp3Player struct {
    Song Song
    //Status string
}

func(m *Mp3Player) Delegate(song Song, trigger chan string) (err error){
    log.SetPrefix("Mp3Player")
    log.Println("Now Into Mp3Player")
    m.Song = song
    go func() {
        trigger <- "current_song"
    }()
    log.Println("I'm into player trigger center")
    select {
        case msg := <-trigger:
            switch msg {
                case "current_song":
                    log.Println( song.Format() )
                    err := m.Play(trigger)
                    if err != nil {
                        // TODO
                    } else {
                        go func(){
                            trigger <- "end_song"
                        }()
                    }
                case "end_song":
                    return nil
                case "loop_song":
                    m.Loop(trigger)
                case "skip_song":
                    return nil
            }
    }
    
    return nil
}

func (m *Mp3Player) Play(trigger chan string) (err error){
    // assumption 3 secs every play though we have no play logic
    time.Sleep(1e9 * 3)
    return nil
}

func (m *Mp3Player) Skip(trigger chan string) (err error){
    go func() {
        trigger <- "skip_song"
    }()
    return nil
}

func (m *Mp3Player) Pause(trigger chan string) (err error){
    return nil
}

func (m *Mp3Player) Resume(trigger chan string) (err error){
    m.Play(trigger)
    return nil
}

func (m *Mp3Player) Fav() (err error){
    return nil
}

func (m *Mp3Player) UnFav() (err error){
    return nil
}

func (m *Mp3Player) Del() (err error){
    return nil
}

func (m *Mp3Player) Loop(trigger chan string) (err error){
    for {
        log.Println("looping...", m.Song.Title)
        m.Play(trigger)
    }
    return nil
}
