package nicehash

type Balance struct {
	Confirmed float64 `json:"balance_confirmed,string"`
	Pending   float64 `json:"balance_pending,string"`
}

func (client *NicehashClient) GetBalance() (Balance, error) {
	version := &struct {
		Result Balance `json:"result"`
	}{}
	params := &Params{Method:"balance", Algo:AlgoTypeMAX, Location:LocationMAX, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&version)
	if err != nil {
		return version.Result, err
	}

	return version.Result, nil
}

