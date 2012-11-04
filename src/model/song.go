package model

type Song struct {
    Url Url
    Title string
    Aid string
    Sid string
    Ssid string
    // TODO Hide these 2 value cause' I don't know why it cause problem now
    //Like int
    //Length int // 
    Subtype string
    Artist string
    Album string
    PublicTime string
    RatingAvg float32
    Company string
    AlbumTitle string
    Picture string
}
