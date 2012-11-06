package model

type FMPlayer interface{
    Delegate(song Song, trigger chan string) (err error)
    Play(trigger chan string) (err error)
    Skip(trigger chan string) (err error)
    Pause(trigger chan string) (err error)
    Resume(trigger chan string) (err error)
    Fav() (err error)
    UnFav() (err error)
    Del() (err error)
    Loop(trigger chan string) (err error)
}
