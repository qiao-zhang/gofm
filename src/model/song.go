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
    Length float64
    Sid string
    Aid string
}

func (s Song) Format() (str string) {
    str = s.Artist + " - " + s.Title + " ("+ s.AlbumTitle +") - " + s.LengthFormat()
    return
}

func (s Song) LengthFormat() (str string) {
    fi := int(s.Length)
    minute, second := fi/60, fi%60
    str = fmt.Sprintf("%0d:%0d", minute, second)
    return
}

func (s Song) IsAdvertisement() (is_ad bool) {
    is_ad = false
    if s.Subtype == "T" {
        is_ad = true
    }
    return
}
