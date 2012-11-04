package model

import (
    "testing"
)

func testInit(t *testing.T) {
    name := "华语"
    intro := "简介"
    song_num := 1024
    c := &Creator{"doubanfm", "http://douban.fm"}
    ch := NewChannel(name, intro, song_num, c)

    if ch.Name != name {
        t.Errorf("ch.Name != name: ", ch.Name, name)
    }
}

