package pixiv

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Transport authorize the requests automatically
// with RefreshToken or Username and Password
type Transport struct {
	Base        http.RoundTripper
	ExpiryDelta time.Duration
	ExpireAt    time.Time
	AuthURL,
	AccessToken,
	RefreshToken,
	Username,
	Password string
	api *AppAPI
}

// RoundTrip implements http.RoundTripper
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() != t.AuthURL {
		if t.AccessToken == "" || t.TokenExpired() {
			_, err := t.Auth(req.Context())
			if err != nil {
				return nil, err
			}
		}
		req.Header.Set("Authorization", "Bearer "+t.AccessToken)
	}

	return t.Base.RoundTrip(req)
}

// TokenExpired checks if the token has expired
func (t *Transport) TokenExpired() bool {
	if t.ExpireAt.IsZero() {
		return false
	}
	return time.Until(t.ExpireAt) < t.ExpiryDelta
}

// Auth get the access_token with refresh_token or username and password.
func (t *Transport) Auth(ctx context.Context) (*RespAuth, error) {
	f := url.Values{
		"client_id":      {t.api.ClientID},
		"client_secret":  {t.api.ClientSecret},
		"device_token":   {t.api.DeviceToken},
		"get_secure_url": {"true"},
		"include_policy": {"true"},
	}
	if t.RefreshToken != "" {
		f.Set("grant_type", "refresh_token")
		f.Set("refresh_token", t.RefreshToken)
	} else if t.Username != "" && t.Password != "" {
		f.Set("grant_type", "password")
		f.Set("username", t.Username)
		f.Set("password", t.Password)
	} else {
		return nil, errors.New("pixiv: refresh_token or username and password not set")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", t.AuthURL, strings.NewReader(f.Encode()))
	t.api.setHeaders(req)
	req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	resp, err := t.Base.RoundTrip(req)
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
		t.AccessToken = r.Response.AccessToken
		if r.Response.ExpiresIn != 0 {
			t.ExpireAt = time.Now().Add(time.Duration(r.Response.ExpiresIn) * time.Second)
		}
		return r, nil
	}
	rerr := &ErrAuth{}
	if rerr.HasError {
		json.Unmarshal(b, rerr)
		return nil, rerr
	}
	return nil, errors.New("pixiv auth: " + string(b))
}
