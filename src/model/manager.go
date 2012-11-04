package model

type Manager interface {
    Channel() (*Channel)
    Playlist() (*Playlist)
    ProgressInPlaylist() (int)
    Player() (FMPlayer)
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

func (m *manager) SetChannalByName(name string) {
    channel := Channel{}
    channel.GetByName(name)
    m.channel = &channel
}

func (m *manager) SetChannelById(id int) {
    channel := Channel{}
    channel.GetById(id)
    m.channel = &channel
}

func (m *manager) CurrentSong() (song *Song) {
    song = &m.playlist.Song[m.progressInPlaylist]
    return
}

func (m *manager) UpdatePlaylist() {
    var typ string
    if m.playlist == nil || len(m.playlist.Song)==0 {
        typ = "n"
    } else if len(m.playlist.Song) == m.progressInPlaylist {
        typ = "p"
    }
    m.playlist.FetchChannel(m.channel.Id, typ)
}

func (m *manager) SetPlayer(p FMPlayer) {
    m.player = p
}
