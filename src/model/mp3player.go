package model

import (
    "time"
)

// Mp3Player 的关注焦点在正确的播放传入的mp3文件
// 它与Manager需要的交互是发送当前的播放状态
type Mp3Player struct {
    Status string
    Source Url
}

func (m *Mp3Player) Play(source Url, trigger chan string) (err error){
    // Hey let's play, every time call play, we need to print song's info in console,
    // so we have to send current_song channel msg to trigger.
    m.Status = "current_song"
    m.Source = source
    trigger <- m.Status

    // assumption 10 secs every play though we have no play logic
    time.Sleep(1e9)
    return nil
}

func (m *Mp3Player) Skip(trigger chan string) (err error){
    m.Status = "skip_song"
    trigger <- m.Status
    return nil
}

func (m *Mp3Player) Pause(trigger chan string) (err error){
    m.Status = "pause_song"
    trigger <- m.Status
    return nil
}

func (m *Mp3Player) Resume(trigger chan string) (err error){
    m.Play(m.Source, trigger)
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
    m.Status = "loop_song"
    trigger <- m.Status
    return nil
}
