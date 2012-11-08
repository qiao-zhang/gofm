package model

import (
    //"testing"
    "fmt"
)

func ExampleFormat() {
    s := new(Song)
    s.Artist = "TestArtist"
    s.Title = "TestTitle"
    s.AlbumTitle = "TestAlbumTitle"
    s.Length = 61

    f := s.Format()
    fmt.Println(f)
    // Output: TestArtist - TestTitle (TestAlbumTitle) - 1:1
}

func ExampleLengthFormat() {
    s := new(Song)
    s.Length = 10.1
    fmt.Println( s.LengthFormat() )

    s.Length = 70
    fmt.Println( s.LengthFormat() )
    // Output: 
    // 0:10
    // 1:10
}

func ExampleIsAdvertisement() {
    s := new(Song)
    s.Subtype = "T"
    fmt.Println( s.IsAdvertisement() )
    s.Subtype = ""
    fmt.Println( s.IsAdvertisement() )
    // Output:
    // true
    // false
}

