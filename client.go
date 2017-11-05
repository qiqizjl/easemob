package easemob

type Client struct {
	ClientID     string
	ClientSecert string
	BaseURI      string
	AdminToken   AdminToken
}

//NewClient 获得一个新的客户端类
func NewClient(ClientID, ClientSecert, OrgName, AppName string) (*Client, error) {
	client := &Client{
		ClientID:     ClientID,
		ClientSecert: ClientSecert,
	}
	client.BaseURI = getBaseURI(OrgName, AppName)
	err := client.getAccessToken()
	if err != nil {
		return nil, err
	}
	return client, err
}

func getBaseURI(OrgName, AppName string) string {
	return "https://a1.easemob.com/" + OrgName + "/" + AppName
}
