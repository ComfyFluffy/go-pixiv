package pixiv

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const oauthURL = "https://oauth.secure.pixiv.net/auth/token"

type respOAuth struct{}

func (api *PixivAppAPI) AuthWithToken(token string) error {
	return api.postOAuth(url.Values{
		"client_id":      {api.ClientID},
		"client_secret":  {api.ClientSecret},
		"device_token":   {api.DeviceToken},
		"get_secure_url": {"true"},
		"include_policy": {"true"},

		"grant_type":    {"refresh_token"},
		"refresh_token": {token},
	})
}

func (api *PixivAppAPI) AuthWithPassword(username, password string) error {
	return api.postOAuth(url.Values{
		"client_id":      {api.ClientID},
		"client_secret":  {api.ClientSecret},
		"device_token":   {api.DeviceToken},
		"get_secure_url": {"true"},
		"include_policy": {"true"},

		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
	})
}

func (api *PixivAppAPI) postOAuth(data url.Values) error {
	req, _ := http.NewRequest("POST", api.OAuthURL,
		strings.NewReader(data.Encode()))
	req.Header = api.HeaderOAuth.Clone()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
	return nil
}
