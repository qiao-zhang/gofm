package model

import (
    "testing"
)

func testInit(t *testing.T) {
    id := 1
    name := "华语"
    intro := "简介"
    song_num := 1024
    hot_songs := []string{"1", "2"}
    c := &Creator{"doubanfm", "http://douban.fm"}
    
    ch := NewChannel(id, name, intro, song_num, hot_songs, c)

    if ch.Name != name {
        t.Errorf("ch.Name != name: ", ch.Name, name)
    }
}

func TestFetchHotChannels(t *testing.T) {

    c := new(Channel)
    total, hot_channels := c.FetchHotChannels()

    if total == 0 {
        t.Error("fetching error")
    }

    if len(hot_channels) == 0 {
        t.Errorf("len(hot_channels) = 0")
    }
}

func TestFetchChannelInfo(t *testing.T) {

    c := new(Channel)
    channel := c.FetchChannelInfo("1")

    if channel.Name != "华语" {
        t.Errorf("channel.Name != 华语", channel.Name)
    }
}
