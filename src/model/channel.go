package model


type Channel struct {
    Id int
    Name string
    Intro string
    SongNum int
    Creator *Creator
}

func NewChannel(id int, name, intro string, song_num int, creator *Creator) *Channel {
    return &Channel{id, name, intro, song_num, creator}
}

func (c *Channel) GetById(id int) {
    // TODO
}


func (c *Channel) GetByName(name string) {
    // TODO
}
