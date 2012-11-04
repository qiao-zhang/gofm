package model

import (
    //"log"
    //"strconv"
    "fmt"
)


type Song struct {
    Picture string
    AlbumTitle string
    Company string
    Rating_avg float32
    PublicTime string
    Ssid string
    Album string
    //Like int
    Artist string
    Url Url
    Title string
    Subtype string
    //Length string
    Sid string
    Aid string
}

func (s Song) Format() (str string) {
    str = s.Artist + " - " + s.Title + " ("+ s.AlbumTitle +")" + s.LengthFormat()
    return
}

func (s Song) LengthFormat() (str string) {
    //fl, _ := s.Length.(float64)
    //fi := int(fl)
    fi := 300
    minute, second := fi/60, fi%60
    str = fmt.Sprintf("%0d:%0d", minute, second)
    return
}
