package easemob

type AdminToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Application string `json:"application"`
}

//获得管理员Token
func (client *Client) getAccessToken() {

}
