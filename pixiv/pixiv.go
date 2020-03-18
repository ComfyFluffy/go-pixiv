package pixiv

import (
	"net/http"
	"time"
)

const (
	clientID     = "KzEZED7aC0vird8jWyHM38mXjNTY"
	clientSecret = "W9JZoJe00qPvJsiyCGT3CCtC6ZUtdpKpzMbNlUGP"
	hashSecret   = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	deviceToken  = "ec731472f8db58afe8588cbba92d5846"
	baseURL      = "https://app-api.pixiv.net"
	imageBaseURL = "https://i.pximg.net"
	timeOut      = 15 * time.Second
)

var baseHeader = http.Header{
	"User-Agent":      {"PixivIOSApp/7.6.7 (iOS 12.3.1; iPhone7,2)"},
	"App-Os":          {"ios"},
	"App-Version":     {"7.6.7"},
	"Accept":          {"*/*"},
	"Accept-Encoding": {"br, gzip, deflate"},
	"Accept-Language": {"en-us"},
}

type PixivAppAPI struct {
	ClientID,
	ClientSecret,
	HashSecret,
	BaseURL,
	OAuthURL,
	ImageBaseURL,
	DeviceToken string
	Client       *http.Client
	HeaderOAuth  http.Header
	HeaderAppAPI http.Header
}

func New() *PixivAppAPI {
	client := &http.Client{
		Timeout: timeOut,
	}
	return NewWithClient(client)
}
func NewWithClient(client *http.Client) *PixivAppAPI {
	// headerOAuth.Set("Accept", "*/*")
	// headerOAuth.Set("Accept-Encoding", "br, gzip, deflate")
	return &PixivAppAPI{
		Client:       client,
		BaseURL:      baseURL,
		OAuthURL:     oauthURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HashSecret:   hashSecret,
		HeaderAppAPI: baseHeader.Clone(),
		HeaderOAuth:  baseHeader.Clone(),
		DeviceToken:  deviceToken,
	}
}

// func (api *PixivAppAPI) baseRequest() *http.Request {
// 	req, _ := http.NewRequest("GET")
// 	return req
// }
