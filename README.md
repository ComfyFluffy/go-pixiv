# go-pixiv

go-pixiv is a Go client for AppAPI of Pixiv.

## Features

* AppAPI
  * Auth
  * User
    * `Detail`
    * `Illusts`
    * `Novels`
    * `BookmarkedIllusts`
    * `BookmarkedNovels`
    * `Followings`
    * `Recommended`
    * `IllustBookmarkTags`
    * `NovelBookmarkTags`
  * Illust
    * `AddBookmark`
    * `DeleteBookmark`
    * `AddHistory`
    * `Comments`
    * `Detail`
    * `Related`
    * `NewFromFollowings`
    * `NewFromAll`
    * `NewFromMyPixiv`
    * `UgoiraMetadata`
    * `RecommendedIllusts`
    * `RecommendedManga`
    * `Ranking`
  * Novel
    * `AddBookmark`
    * `DeleteBookmark`
    * `AddHistory`
    * `Text`
    * `Comments`
    * `Detail`
    * `Recommended`
    * `Ranking`
  * Comment
    * `RepliesIllust`
    * `RepliesNovel`
    * `AddToIllust`
    * `AddToNovel`
    * `DeleteFromIllust`
    * `DeleteFromNovel`
  * Search
    * `IllustTrendingTags`
    * `NovelTrendingTags`
    * `Illusts`
    * `PopularIllustsPreview`
    * `Novels`
    * `PopularNovelsPreview`
    * `TagsStartWith`
    * `Users`
* Proxy support
  * HTTP
  * SOCKS5
* TODO
  * [] Bypass SNI filtering with DoH (Bypass GFW)

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

    r2, err := api.Search.Illusts("風景", &pixiv.SearchQuery{
        SearchTarget: pixiv.STExactMatchTags,
        Sort: pixiv.SDateDesc,
    })
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("%+v", r2)
    log.Print(r2.NextIllusts())
}
```
