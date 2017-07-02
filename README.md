# Instagram Go Scraper 
[![GoDoc](https://godoc.org/github.com/Vorkytaka/instagram-go-scraper/instagram?status.svg)](https://godoc.org/github.com/Vorkytaka/instagram-go-scraper/instagram)
[![Build Status](https://travis-ci.org/Vorkytaka/instagram-go-scraper.svg?branch=master)](https://travis-ci.org/Vorkytaka/instagram-go-scraper.svg?branch=master)

Instagram Scraper for Golang.

A package that helps you with requesting to Instagram without a key.

### Installation:
Install:
```
$ go get -u github.com/Vorkytaka/instagram-go-scraper/instagram
```
Import:
```go
import "github.com/Vorkytaka/instagram-go-scraper/instagram"
```

### List of functions:
After import you can use following functions:
```go
// Get account info
account, err := instagram.GetAccountByUsername("username")

// Get media info
media, err := instagram.GetMediaByCode("code")
media, err := instagram.GetMediaByURL("https://instagram.com/p/code")

// Get slice of account media
media, err := instagram.GetAccountMedia("username", limit)
// or slice of all account media
media, err := instagram.GetAllAccountMedia("username")

// Get slice of location last media
media, err := instagram.GetLocationMedia("location_id", limit)
// Get array[9] of location top media
media, err := instagram.GetLocationTopMedia("location_id")

// Get location info
location, err := instagram.GetLocationByID("location_id")

// Get slice of tag last media
media, err := instagram.GetTagMedia("tag", limit)
// Get array[9] of tag top media
media, err := instagram.GetLocationTopMedia("tag")

// Search for users (return slice of Accounts)
users, err := instagram.SearchForUsers("username")
```

You'll get `err != nil` if request return 404 or if there a parsing error.

### About media updates:

Media can have one of 3 types:
* `TypeImage`
* `TypeVideo`
* `TypeCarousel`

From v0.2.03 media has MediaList field for collection of media.

If media is carousel, then there will be a list of all media.

If media is image or video, then there will be only one media, but also,
for backward compatibility, media url will be in `media.MediaURL` field.

### Update data:

You can update media by call `Update` method:
```go
media, err := instagram.GetMediaByCode("code")
err := media.Update()
if err != nil {
    // media didn't update
}
```
Same with account:
```go
account, err := instagram.GetAccountByUsername("username")
err := account.Update()
if err != nil {

}
```
