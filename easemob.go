package easemob

type Client struct {
	ClientID     string
	ClientSecert string
	BaseURI      string
	AdminToken   AdminToken
}

//NewClient 获得一个新的客户端类
func NewClient(ClientID, ClientSecert, OrgName, AppName string) *Client {
	client := &Client{
		ClientID,
		ClientSecert,
	}
	client.BaseURI = getBaseURI(OrgName, AppName)
	return client
}

func getBaseURI(OrgName, AppName string) string {
	return "https://a1.easemob.com/" + OrgName + "/" + AppName
}
