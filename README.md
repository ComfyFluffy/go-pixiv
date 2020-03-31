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
    api.SetRefreshToken("xxx")
    api.SetProxy("socks5://127.0.0.1")
    r, err := api.User.Detail(123, nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%+v", r)
}
```
