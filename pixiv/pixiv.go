package pixiv

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

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

// Proxy setting errors
var (
	ErrSetProxyUnsupportedTransport = errors.New("pixiv: can only set proxy for *http.Transport")
	ErrSetProxyUnsupportedProtocol  = errors.New("pixiv: unsupported proxy protocol")
)

type service struct {
	api *AppAPI
}

// AppAPI defines the Pixiv App-API client with config.
type AppAPI struct {
	ClientID,
	ClientSecret,
	HashSecret,
	BaseURL,
	ImageBaseURL,
	DeviceToken string
	BaseHeader http.Header

	Client *http.Client // *http.Client with *Transport that can authorize requests automatically

	service *service
	User    *UserService
}

func (api *AppAPI) transportAuth() *Transport {
	return api.Client.Transport.(*Transport)
}

// New returns new PixivAppAPI with http.DefaultClient
func New() *AppAPI {
	return NewWithClient(&http.Client{Timeout: timeOut, Transport: http.DefaultTransport.(*http.Transport).Clone()})
}

// NewWithClient returns new PixivAppAPI with the given http.Client.
func NewWithClient(client *http.Client) *AppAPI {
	api := &AppAPI{
		BaseURL:      baseURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HashSecret:   hashSecret,
		DeviceToken:  deviceToken,
		BaseHeader:   baseHeader.Clone(),
		Client:       client,
	}

	api.service = &service{api: api}
	api.User = (*UserService)(api.service)

	client.Transport = &Transport{
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
	default:
		return ErrSetProxyUnsupportedProtocol
	}
	return nil
}

// SetUser sets the username and password for auth.
func (api *AppAPI) SetUser(username, password string) {
	tr := api.transportAuth()
	tr.Username = username
	tr.Password = password
	tr.RefreshToken = ""
}

// SetRefreshToken sets the refresh_token for auth.
func (api *AppAPI) SetRefreshToken(token string) {
	tr := api.transportAuth()
	tr.RefreshToken = token
	tr.Username = ""
	tr.Password = ""
}

// SetLanguage sets Accept-Language header to the given language.
// This affects the language of tag translations and messages.
func (api *AppAPI) SetLanguage(lang string) {
	api.BaseHeader["Accept-Language"] = []string{lang}
}

// Auth do the auth with given username and password or refresh_token.
func (api *AppAPI) Auth() (*RespAuth, error) {
	return api.transportAuth().Auth(context.Background())
}

func (api *AppAPI) setHeaders(req *http.Request) {
	req.Header = api.BaseHeader.Clone()
	nows := time.Now().Format(time.RFC3339)
	req.Header["X-Client-Time"] = []string{nows}
	x := md5.Sum([]byte(nows + api.HashSecret))
	req.Header["X-Client-Hash"] = []string{hex.EncodeToString(x[:])}
}

// NewRequest sets headers and body of a new request with given method, url and form.
func (api *AppAPI) NewRequest(method, url string, data url.Values) (*http.Request, error) {
	var buf io.Reader
	if data != nil {
		buf = strings.NewReader(data.Encode())
	}
	req, err := http.NewRequest(method, url, buf)
	api.setHeaders(req)
	if err != nil {
		return nil, err
	}
	if data != nil {
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	}
	return req, err
}

// Receive sends the request and decode the response into successV or errorV.
// If the status code is 2XX, the response will be decode into successV.
// Otherwise, it will be decode into errorV.
func (api *AppAPI) Receive(req *http.Request, successV interface{}, errorV interface{}) (bool, *http.Response, error) {
	resp, err := api.Client.Do(req)
	if err != nil {
		return false, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 300 && resp.StatusCode >= 200 {
		if successV != nil {
			dec := json.NewDecoder(resp.Body)
			err = dec.Decode(successV)
			if err != nil {
				return false, nil, err
			}
		}
		return true, resp, err
	}
	if errorV != nil {
		dec := json.NewDecoder(resp.Body)
		err = dec.Decode(errorV)
		if err != nil {
			return false, nil, err
		}
	}
	return false, resp, nil
}

func (api *AppAPI) get(purl string, query url.Values, v interface{}) (*http.Response, error) {
	req, err := api.NewRequest("GET", api.BaseURL+purl, nil)
	if err != nil {
		return nil, err
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	rerr := &ErrAppAPI{}
	ok, resp, err := api.Receive(req, v, rerr)
	if err != nil {
		return nil, err
	}
	if !ok {
		rerr.response = resp
		return nil, rerr
	}

	return resp, nil
}