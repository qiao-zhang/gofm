package model

import (
    "testing"
)

func TestGetManagerInstance(t *testing.T) {
    m1 := GetManagerInstance()
    m2 := GetManagerInstance()
    if m1 != m2 {
        t.Errorf("two instance are not the same")
    }
}

func TestChooseChannel(t *testing.T) {
    m := GetManagerInstance()
    m.ChooseChannel("1")
    if m.Playlist() == nil {
        t.Errorf("After chooseChannel, playlist should be update!")
    }
}
