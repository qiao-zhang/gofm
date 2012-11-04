package model


type Channel struct {
    Name string
    Intro string
    SongNum int
    Creator *Creator
}

func NewChannel(name, intro string, song_num int, creator *Creator) *Channel {
    return &Channel{name, intro, song_num, creator}
}
