package model

import (
    "testing"
)

func TestFetchChannalPlaylist(t *testing.T) {
    playlist := new(Playlist)
    playlist.FetchChannel(1, "n")

    if playlist.R != 0 {
        t.Errorf( playlist.Err )
    }

    if len(playlist.Song) == 0 {
        t.Errorf( "%v", playlist.Song )
    }
}
