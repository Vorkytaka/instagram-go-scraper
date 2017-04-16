# Instagram Go Scraper
Instagram Scraper for Golang.

A package that helps you with requesting to Instagram without a key.

### Installation:
Install:
```
$ go get -u github.com/SolidlSnake/instagram-go-scraper/instagram
```
Import:
```go
import "github.com/SolidlSnake/instagram-go-scraper/instagram"
```

### List of functions:
After import you can use following functions:
```go
// Get account info
account, err := instagram.GetAccoutByUsername("username")

// Get media info
media, err := instagram.GetMediaByCode("code")
media, err := instagram.GetMediaByUrl("https://instagram.com/p/code")

// Get slice of account media
media, err := instagram.GetAccountMedia("username", limit)
// or slice of all account media
media, err := instagram.GetAllAccountMedia("username")

// Get slice of location last media
media, err := instagram.GetLocationMedia("location_id", limit)
// Get array[9] of location top media
media, err := instagram.GetLocationTopMedia("location_id")

// Get location info
location, err := instagram.GetLocationById("location_id")

// Get slice of tag last media
media, err := instagram.GetTagMedia("tag", limit)
// Get array[9] of tag top media
media, err := instagram.GetLocationTopMedia("tag")

```

You'll get `err != nil` if request return 404 or if there a parsing error.