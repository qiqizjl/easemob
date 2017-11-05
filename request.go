package easemob

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//发送请求
func (client *Client) doRequest(uri string, body io.Reader, method string, needAdminToken bool) (string, error) {
	//请求地址
	requetURI := client.BaseURI + uri
	//初始化一个Client
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, requetURI, body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if needAdminToken == true {
		req.Header.Set("Authorization", "Bearer "+client.GetToken())
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized && uri != "/token" {
			//401表示token过期，重新刷新下请求
			err := client.getAccessToken()
			if err != nil {
				return "", err
			}
			return client.doRequest(uri, body, method, needAdminToken)
		}
		return string(result), fmt.Errorf(res.Status)
	}
	return string(result), nil
}
