package easemob

func (client *Client) doRequest(uri string, body io.Reader, method string) (string, error)
	//请求地址
	requetURI := client.BaseURI + uri
	//初始化一个Client
	client := &http.Client{}
	req, err := http.NewRequest(method, requetURI, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if c.adminToken.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.adminToken.AccessToken)
	}

}
