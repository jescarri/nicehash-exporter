package nicehash

type Version struct {
	ApiVersion  string `json:"api_version"`
}

func (client *NicehashClient) GetVersion() (string, error) {
	version := &struct {
		Result Version `json:"result"`
	}{}
	_, err := client.sling.New().Get("").ReceiveSuccess(&version)
	if err != nil {
		return version.Result.ApiVersion, err
	}

	return version.Result.ApiVersion, err
}
