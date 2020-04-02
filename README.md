# go-pixiv

go-pixiv is a Go client for AppAPI of Pixiv.

## Features

* AppAPI
  * Auth
  * User
    * Detail
    * Illusts
    * Novels
    * Bookmarks
    * Followings
    * Recommend
    * BookmarkTags
  * Illust
    * AddBookmark
    * DeleteBookmark
    * AddHistory
    * Comments
    * Detail
    * Related
    * NewFromFollowings
    * NewFromAll
    * NewFromMyPixiv
    * UgoiraMetadata
  * Novel
    * AddBookmark
    * DeleteBookmark
    * AddHistory
    * Text
    * Comments
    * Detail
  * Comment
    * Replies
  * More in progress...
* Proxy support
  * HTTP
  * SOCKS5

## Install

`go get github.com/WOo0W/go-pixiv`

## Example

```go
package main

import (
    "log"
    "github.com/WOo0W/go-pixiv/pixiv"
)

func main() {
    api := pixiv.New()

    api.SetUser("xxx@xxx.com", "password")
    // or api.SetRefreshToken("xxx")

    api.SetProxy("socks5://127.0.0.1:1080")

    r, err := api.User.Detail(123, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%+v", r)
    api.Illust.AddBookmark(123, pixiv.RPublic, nil)
    api.Illust.AddBookmark(456, pixiv.RPrivate,
        &pixiv.AddBookmarkOptions{
        Tags: []string{"風景"},
    })
}
```
