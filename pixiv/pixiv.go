package pixiv

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	clientID     = "KzEZED7aC0vird8jWyHM38mXjNTY"
	clientSecret = "W9JZoJe00qPvJsiyCGT3CCtC6ZUtdpKpzMbNlUGP"
	hashSecret   = "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"
	deviceToken  = "ec731472f8db58afe8588cbba92d5846"
	baseURL      = "https://app-api.pixiv.net"
	authURL      = "https://oauth.secure.pixiv.net/auth/token"
	timeOut      = 15 * time.Second
	expiryDelta  = 30 * time.Second
)

var baseHeader = http.Header{
	"User-Agent":     {"PixivIOSApp/7.8.30 (iOS 12.4.5; iPhone7,2)"},
	"App-OS":         {"ios"},
	"App-OS-Version": {"12.4.6"},
	"App-Version":    {"7.8.30"},
	"Accept":         {"*/*"},
	// "Accept-Encoding": {"br, gzip, deflate"},
	"Accept-Language": {"en-us"},
}

type service struct {
	api *AppAPI
}

// AppAPI defines the Pixiv App-API client with config.
type AppAPI struct {
	ClientID,
	ClientSecret,
	HashSecret,
	BaseURL,
	DeviceToken string
	BaseHeader http.Header

	AuthURL,
	Username,
	Password,
	RefreshToken,
	AccessToken string
	UserID           int
	TokenExpireAt    time.Time
	TokenExpiryDelta time.Duration

	// Contains details of login user.
	AuthResponse *RespAuth

	Client *http.Client // *http.Client with *Transport that can authorize requests automatically

	service *service

	User    *UserService
	Illust  *IllustService
	Novel   *NovelService
	Comment *CommentService
	Search  *SearchService
}

// New returns new PixivAppAPI with http.DefaultClient
func New() *AppAPI {
	return NewWithClient(&http.Client{Timeout: timeOut, Transport: &http.Transport{}})
}

// NewWithClient returns new PixivAppAPI with the given http.Client.
func NewWithClient(client *http.Client) *AppAPI {
	api := &AppAPI{
		BaseURL:      baseURL,
		AuthURL:      authURL,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HashSecret:   hashSecret,
		DeviceToken:  deviceToken,
		BaseHeader:   baseHeader.Clone(),
		Client:       client,
	}

	api.service = &service{api: api}
	api.User = (*UserService)(api.service)
	api.Illust = (*IllustService)(api.service)
	api.Novel = (*NovelService)(api.service)
	api.Comment = (*CommentService)(api.service)
	api.Search = (*SearchService)(api.service)

	return api
}

// SetUser sets the username and password for auth.
func (api *AppAPI) SetUser(username, password string) {
	api.Username = username
	api.Password = password
	api.RefreshToken = ""
}

// SetRefreshToken sets the refresh_token for auth.
func (api *AppAPI) SetRefreshToken(token string) {
	api.RefreshToken = token
	api.Username = ""
	api.Password = ""
}

// SetLanguage sets Accept-Language header to the given languages.
// This affects the language of tag translations and messages.
func (api *AppAPI) SetLanguage(language string) {
	api.BaseHeader["Accept-Language"] = []string{language}
}

// SetHeaders sets the header of req with BaseHeader
// and adds X-Client-Time & X-Client-Hash headers.
func (api *AppAPI) SetHeaders(req *http.Request) {
	req.Header = api.BaseHeader.Clone()
	nows := time.Now().Format(time.RFC3339)
	req.Header["X-Client-Time"] = []string{nows}
	x := md5.Sum([]byte(nows + api.HashSecret))
	req.Header["X-Client-Hash"] = []string{hex.EncodeToString(x[:])}
}

func readerFromForm(data url.Values) io.Reader {
	if data != nil {
		return strings.NewReader(data.Encode())
	}
	return nil
}

// NewAuthorizedRequest sets auth and other headers and body of a new request
// with given method, url and form data.
func (api *AppAPI) NewAuthorizedRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if api.AccessToken == "" || api.TokenExpired() {
		_, err := api.ForceAuth()
		if err != nil {
			return nil, err
		}
	}

	api.SetHeaders(req)
	req.Header["Authorization"] = []string{"Bearer " + api.AccessToken}
	if body != nil {
		req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	}

	return req, nil
}

// NewPximgRequest sets base headers and sets Referer to "https://app-api.pixiv.net/"
func (api *AppAPI) NewPximgRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	api.SetHeaders(req)
	req.Header["Referer"] = []string{"https://app-api.pixiv.net/"}
	return req, nil
}

// receive sends the request and decode the response into successV or errorV.
// If the status code is 2XX, the response will be decode into successV.
// Otherwise, it will be decode into errorV.
func (api *AppAPI) receive(req *http.Request, successV interface{}, errorV interface{}) (bool, *http.Response, error) {
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
		return true, resp, nil
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

func (api *AppAPI) withAppAPIErrors(req *http.Request, v interface{}) (*http.Response, error) {
	rerr := &ErrAppAPI{}
	ok, resp, err := api.receive(req, v, rerr)
	if err != nil {
		return nil, err
	}
	if !ok {
		rerr.response = resp
		return nil, rerr
	}
	return resp, nil
}

func (api *AppAPI) get(r interface{}, urls string, query url.Values) error {
	req, err := api.NewAuthorizedRequest("GET", urls, nil)
	if err != nil {
		return err
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	_, err = api.withAppAPIErrors(req, r)
	return err
}

func (api *AppAPI) post(r interface{}, urls string, data url.Values) error {
	req, err := api.NewAuthorizedRequest("POST", urls, readerFromForm(data))
	if err != nil {
		return err
	}

	_, err = api.withAppAPIErrors(req, r)
	return err
}

func (api *AppAPI) getWithValues(r interface{}, urls string, opts interface{}, values url.Values, caller string) error {
	q, err := withOpts(opts, values, caller)
	if err != nil {
		return err
	}

	return api.get(r, urls, q)
}

func (api *AppAPI) postWithValues(r interface{}, urls string, opts interface{}, values url.Values, caller string) error {
	body, err := withOpts(opts, values, caller)
	if err != nil {
		return err
	}

	return api.post(r, urls, body)
}
