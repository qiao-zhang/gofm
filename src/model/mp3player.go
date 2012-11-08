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
}

func(m *Mp3Player) Delegate(song Song, trigger chan string) (err error){
    log.SetPrefix("Mp3Player ")
    m.Song = song
    for {
        select {
            case msg := <-trigger:
                switch msg {
                    case "current_song":
                        log.Println( song.Format() )
                        go func(){
                            err := m.Play(trigger)
                            if err != nil {
                                // TODO
                            } else {

                            }
                        }()
                    case "end_song":
                        log.Println("end_song", song.Format())
                        goto done // I like goto, whatever.
                    case "loop_song":
                        m.Loop(trigger)
                    case "skip_song":
                        // TODO
                        log.Println("skip signal...")
                        break
                    case "pause_song":
                        // TODO
                        log.Println("recieved pause signal...")
                    default:
                        log.Println("waiting...")
                }
        }
    }
    done:
    log.Println("Delegate done")
    return nil
}


func (m *Mp3Player) Play(trigger chan string) (err error){
    // assumption 3 secs every play though we have no play logic
    for i:=0; i < 5; i++ {
        log.Println("playing ... ")
        time.Sleep( time.Second )
    }
    go func(){
        trigger <- "end_song"
    }()
    return nil
}

func (m *Mp3Player) Skip(trigger chan string) (err error){
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
    go func(){
        // TODO Fetch url of trash one song
    }()
    
    return nil
}

func (m *Mp3Player) Loop(trigger chan string) (err error){
    // TODO current bug: will loop next song
    for {
        log.Println("looping...", m.Song.Title)
        m.Play(trigger)
    }
    return nil
}
