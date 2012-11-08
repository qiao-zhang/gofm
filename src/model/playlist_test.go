package model

import (
    "testing"
)

func TestFetchChannel(t *testing.T) {
    playlist := new(Playlist)
    playlist.FetchChannel(1, "n")

    if playlist.R != 0 {
        t.Errorf( playlist.Err )
    }

    if len(playlist.Song) == 0 {
        t.Errorf( "%v", playlist.Song )
    }
}

func TestFetchChannelNextSong(t *testing.T) {
    playlist := new(Playlist)
    channel := 1
    typ := "p"
    sid := "1457109"
    playlist.FetchChannelNextSong( channel, typ, sid)

    if len(playlist.Song) == 0 {
        t.Errorf( "%v", playlist)
    }
}

