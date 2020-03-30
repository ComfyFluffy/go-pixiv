package pixiv

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dghubble/sling"
	"golang.org/x/net/proxy"
)

const (
	clientID     = "KzEZED7aC0vird8jWyHM38mXjNTY"
	clientSecret = "W9JZoJe00qPvJsiyCGT3CCtC6ZUtdpKpzMbNlUGP"
	hashSecret   = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	deviceToken  = "ec731472f8db58afe8588cbba92d5846"
	baseURL      = "https://app-api.pixiv.net"
	authURL      = "https://oauth.secure.pixiv.net/auth/token"
	imageBaseURL = "https://i.pximg.net"
	timeOut      = 15 * time.Second
	expiryDelta  = 30 * time.Second
)

var baseHeader = http.Header{
	"User-Agent":     {"PixivIOSApp/7.8.16 (iOS 12.4.5; iPhone7,2)"},
	"App-OS":         {"ios"},
	"App-OS-Version": {"12.4.5"},
	"App-Version":    {"7.8.16"},
	"Accept":         {"*/*"},
	// "Accept-Encoding": {"br, gzip, deflate"},
	"Accept-Language": {"en-us"},
}

var (
	ErrSetProxyUnsupportedTransport = errors.New("pixiv: can only set proxy for *http.Transport")
	ErrSetProxyUnsupportedProtocol  = errors.New("pixiv: unsupported proxy protocol")
)

// AppAPI defines the Pixiv App-API client with config
type AppAPI struct {
	ClientID,
	ClientSecret,
	HashSecret,
	BaseURL,
	ImageBaseURL,
	DeviceToken string
	BaseHeader http.Header

	Sling  *sling.Sling
	Client *http.Client // *http.Client with *Transport that can authorize requests automatically
}

func (api *AppAPI) transportAuth() *TransportAuth {
	return api.Client.Transport.(*TransportAuth)
}

// New returns new PixivAppAPI with http.DefaultClient
func New() *AppAPI {
	return NewWithClient(&http.Client{Timeout: timeOut, Transport: http.DefaultTransport.(*http.Transport).Clone()})
}

// NewWithClient returns new PixivAppAPI with the given http.Client
func NewWithClient(client *http.Client) *AppAPI {
	s := sling.New().Client(client)
	for k, v := range baseHeader {
		s.Set(k, v[0])
	}

	api := &AppAPI{
		Sling:        s,
		BaseURL:      baseURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HashSecret:   hashSecret,
		DeviceToken:  deviceToken,
		BaseHeader:   baseHeader.Clone(),
		Client:       client,
	}

	client.Transport = &TransportAuth{
		Base:        client.Transport,
		AuthURL:     authURL,
		ExpiryDelta: expiryDelta,
		api:         api,
	}

	return api
}

// SetProxy sets the proxy with the given URI.
// Supports SOCKS5 or HTTP proxy.
func (api *AppAPI) SetProxy(p string) error {
	trb := api.transportAuth().Base
	var tr *http.Transport
	if trx, ok := trb.(*http.Transport); ok {
		tr = trx
	} else {
		return ErrSetProxyUnsupportedTransport
	}

	pr, err := url.Parse(p)
	if err != nil {
		return err
	}

	switch strings.ToLower(pr.Scheme) {
	case "http":
		hp := http.ProxyURL(pr)
		tr.Proxy = hp
	case "socks5":
		var spauth *proxy.Auth
		spw, _ := pr.User.Password()
		if spw != "" || pr.User.Username() != "" {
			spauth = &proxy.Auth{User: pr.User.Username(), Password: spw}
		}
		spd, err := proxy.SOCKS5("tcp", pr.Host, spauth, proxy.Direct)
		if err != nil {
			return err
		}
		tr.DialContext = spd.(proxy.ContextDialer).DialContext
	}
	return ErrSetProxyUnsupportedProtocol
}

func (api *AppAPI) SetUser(username, password string) {
	tr := api.transportAuth()
	tr.Username = username
	tr.Password = password
	tr.RefreshToken = ""
}

func (api *AppAPI) SetRefreshToken(token string) {
	tr := api.transportAuth()
	tr.RefreshToken = token
	tr.Username = ""
	tr.Password = ""
}

func (api *AppAPI) Auth() (*RespAuth, error) {
	return api.transportAuth().Auth(context.Background())
}
