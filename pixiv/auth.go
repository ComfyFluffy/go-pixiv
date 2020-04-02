package pixiv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TokenExpired checks if the token has expired
func (api *AppAPI) TokenExpired() bool {
	if api.TokenExpireAt.IsZero() {
		return false
	}
	return time.Until(api.TokenExpireAt) < api.TokenExpiryDelta
}

// ForceAuth gets new access_token with given username and password or refresh_token wether it expires.
func (api *AppAPI) ForceAuth() (*RespAuth, error) {
	f := url.Values{
		"client_id":      {api.ClientID},
		"client_secret":  {api.ClientSecret},
		"device_token":   {api.DeviceToken},
		"get_secure_url": {"true"},
		"include_policy": {"true"},
	}
	if api.RefreshToken != "" {
		f.Set("grant_type", "refresh_token")
		f.Set("refresh_token", api.RefreshToken)
	} else if api.Username != "" && api.Password != "" {
		f.Set("grant_type", "password")
		f.Set("username", api.Username)
		f.Set("password", api.Password)
	} else {
		return nil, errors.New("pixiv: refresh_token or username and password not set")
	}

	req, err := http.NewRequest("POST", api.AuthURL, strings.NewReader(f.Encode()))
	api.setHeaders(req)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 300 && resp.StatusCode >= 200 {
		r := &RespAuth{}
		err := json.Unmarshal(b, r)
		if err != nil {
			return nil, fmt.Errorf("pixiv auth: json decode: %w", err)
		}
		if r.Response.AccessToken == "" {
			return nil, errors.New("pixiv auth: no access_token received")
		}
		api.AccessToken = r.Response.AccessToken
		if r.Response.ExpiresIn != 0 {
			api.TokenExpireAt = time.Now().Add(time.Duration(r.Response.ExpiresIn) * time.Second)
		}
		api.AuthResponse = r
		return r, nil
	}
	rerr := &ErrAuth{}
	if rerr.HasError {
		json.Unmarshal(b, rerr)
		return nil, rerr
	}
	return nil, errors.New("pixiv auth: " + string(b))
}
