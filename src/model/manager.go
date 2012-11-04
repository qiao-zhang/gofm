package model

import (
    //"time"
    "log"
    "strconv"
)

type Manager interface {
    Channel() (*Channel)
    Playlist() (*Playlist)
    ProgressInPlaylist() (int)
    Player() (FMPlayer)
    ChooseChannel(id string)
    Start(trigger chan string)
}

type manager struct{
    channel *Channel
    playlist *Playlist
    progressInPlaylist int
    player FMPlayer
}

var theManager Manager
func GetManagerInstance() Manager {
    if theManager == nil {
        mp3_player := new(Mp3Player)
        theManager = & manager {nil, nil, 0, mp3_player}
    }
    return theManager
}

func (m *manager) Channel() (*Channel) {
    if m.channel == nil {
        m.channel = &Channel{}
        //m.channel.SetById(1)
    }
    return m.channel
}

func (m *manager) Playlist() (*Playlist) {
    return m.playlist
}

func (m manager) ProgressInPlaylist() (int) {
    return m.progressInPlaylist
}

func (m *manager) Player() (FMPlayer) {
    return m.player
}

func (m *manager) ChooseChannel(channel_id string) {
    //go func() {
        //ch := new(Channel)
        //m.channel = ch.FetchChannelInfo(channel_id)
    //}()

    m.playlist = nil
    m.progressInPlaylist = 0
    cid, err := strconv.Atoi(channel_id)
    if err != nil {
        log.Printf("U should type a number id")
        return
    }
    m.UpdatePlaylist(cid)
}

func (m *manager) CurrentSong() (song *Song) {
    song = &m.playlist.Song[m.progressInPlaylist]
    return
}

func (m *manager) UpdatePlaylist(channel_id int) {
    defer func() {
        if r:= recover() ; r != nil {
            log.Printf("cannot update playlist")
        }
    }()
    var typ string
    if m.playlist == nil  || len(m.playlist.Song)==0 {
        typ = "n"
        m.playlist = new(Playlist)
        m.playlist.FetchChannel(channel_id, typ)
    } else {
        typ = "p"
        m.playlist.FetchChannelNextSong(channel_id, typ,
                m.playlist.Song[m.progressInPlaylist].Sid)
    }
}

func (m *manager) SetPlayer(p FMPlayer) {
    m.player = p
}

func (m *manager) Start(trigger chan string)  {
    // init channel, first get into channel 1
    m.ChooseChannel("1")
    channel_id := 1
    if m.playlist != nil {
        // init player, first send current_song
        for {
            playlist := m.playlist
            m.progressInPlaylist = 0
            for i:=0; i<len(playlist.Song) ; i++ {
                m.progressInPlaylist = i
                m.player.Play(playlist.Song[m.progressInPlaylist].Url, trigger)
            }
            m.UpdatePlaylist( channel_id ) // may stop several seconds
        }
    }
}
