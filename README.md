# Instagram Go Scraper
Instagram Scraper for Golang.

A package that helps you with requesting to Instagram without a key.

### Installation:
Install:
```
$ go get -u github.com/SolidlSnake/instagram-go-scraper/instagram
```
Imoort:
```go
import "github.com/SolidlSnake/instagram-go-scraper/instagram"
```

### List of functions:
After import you can use following functions:
```go
// Get account info
account, err := instagram.GetAccountMedia("username")

// Get media info
media, err := instagram.GetMediaByCode("code")
media, err := instagram.GetMediaByUrl("https://instagram.com/p/code")

// Get slice of account media
account, err := instagram.GetAccountMedia("username", limit)
// or slice of all account media
account, err := instagram.GetAllAccountMedia("username")

// Get slice of location last media
account, err := instagram.GetLocationMedia("location_id", limit)
// Get array[9] of location top media
account, err := instagram.GetLocationTopMedia("location_id")

// Get location info
location, err := instagram.GetLocationById("location_id")

```

You'll get `err != nil` if request return 404 or if there a parsing error.