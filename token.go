package easemob

import (
	"bytes"
	"encoding/json"
	"time"
)

type AdminToken struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	Application string    `json:"application"`
	Time        time.Time `json:"-"`
}

//获得管理员Token
func (client *Client) getAccessToken() error {
	requestMap := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     client.ClientID,
		"client_secret": client.ClientSecert,
	}
	b, _ := json.Marshal(requestMap)
	result, err := client.doRequest("/token", bytes.NewBuffer([]byte(b)), "POST", false)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(result), &client.AdminToken)
	client.AdminToken.Time = time.Now().Add(time.Duration(client.AdminToken.ExpiresIn-60*10) * time.Second)
	return nil
}

//GetToken 获得Token
func (client *Client) GetToken() string {
	if client.AdminToken.Time.Unix() < time.Now().Unix() {
		//获取Token
		client.getAccessToken()
	}
	return client.AdminToken.AccessToken
}
