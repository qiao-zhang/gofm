package model

type FMPlayer interface{
    Play() (err error)
    Skip() (err error)
    Pause() (err error)
    Resume() (err error)
    Fav() (err error)
    UnFav() (err error)
    Del() (err error)
    Loop() (err error)
}
