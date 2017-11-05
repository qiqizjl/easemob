package easemob

import (
	"bytes"
	"encoding/json"
	"strings"
)

//Message2 发送消息
func (client *Client) Message2(body string) (string, error) {
	return client.doRequest("/messages", strings.NewReader(string(body)), "POST", true)
}

//Message 发送消息
func (client *Client) Message(targetType string, target []string, msg interface{}, from string) (string, error) {
	postBody := map[string]interface{}{
		"target_type": targetType,
		"target":      target,
		"msg":         msg,
	}
	if from != "" {
		postBody["from"] = from
	}
	str, _ := json.Marshal(postBody)
	return client.doRequest("/messages", bytes.NewReader(str), "POST", true)
}
