package nicehash

type Orders struct {
	Id            uint64 `json:"id"`
	Type          OrderType `json:"type"`
	Algo          AlgoType `json:"algo"`
	Price         float64 `json:"price,string"`
	Alive         bool `json:"alive"`
	LimitSpeed    float64 `json:"limit_speed,string"`
	AcceptedSpeed float64 `json:"accepted_speed,string"`
	Workers       uint64 `json:"workers"`
}

func (client *NicehashClient) GetOrders(algo AlgoType, location Location) ([]Orders, error) {
	stats := &struct {
		Result struct {
			       Orders []Orders `json:"orders"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.get", Algo:algo, Location:location}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Orders, err
	}

	return stats.Result.Orders, nil
}

type MyOrders struct {
	Id            uint64 `json:"id"`
	Type          OrderType `json:"type"`
	Algo          AlgoType `json:"algo"`
	Price         float64 `json:"price,string"`
	BtcAvail      float64 `json:"btc_avail,string"`
	BtcPaid       float64 `json:"btc_paid,string"`
	PoolHost      string `json:"pool_host"`
	PoolPort      uint16 `json:"pool_port"`
	PoolUser      string `json:"pool_user"`
	PoolPass      string `json:"pool_pass"`
	Alive         bool `json:"alive"`
	LimitSpeed    float64 `json:"limit_speed,string"`
	AcceptedSpeed float64 `json:"accepted_speed,string"`
	Workers       uint64 `json:"workers"`
	End           uint64 `json:"end"`
}

func (client *NicehashClient) GetMyOrders(algo AlgoType, location Location) ([]MyOrders, error) {
	stats := &struct {
		Result struct {
			       Orders []MyOrders `json:"orders"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.get", Algo:algo, Location:location, My:true, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Orders, err
	}

	return stats.Result.Orders, nil
}

type NewOrder struct {
	Algo       AlgoType `json:"algo" url:"algo"`
	Price      float64 `json:"price,string" url:"price"`
	Amount     float64 `json:"amount,string" url:"amount"`
	PoolHost   string `json:"pool_host" url:"pool_host"`
	PoolPort   uint16 `json:"pool_port" url:"pool_port"`
	PoolUser   string `json:"pool_user" url:"pool_user"`
	PoolPass   string `json:"pool_pass" url:"pool_pass"`
	Alive      bool `json:"alive" url:"alive"`
	LimitSpeed float64 `json:"limit,string" url:"limit,omitempty"`
	Code       string `json:"code" url:"code,omitempty"`
}

func (client *NicehashClient) OrderCreate(order NewOrder) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.create", Algo:AlgoTypeMAX, Location:LocationMAX, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).QueryStruct(order).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}

func (client *NicehashClient) OrderRefill(algo AlgoType, location Location, order uint, amount float64) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.refill", Order:order, Algo:algo, Location:location, Amount:amount, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}

func (client *NicehashClient) OrderRemove(algo AlgoType, location Location, order uint) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.remove", Order:order, Algo:algo, Location:location, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}

func (client *NicehashClient) OrderSetPrice(algo AlgoType, location Location, order uint, price float32) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.set.price", Algo:algo, Location:location, Order:order, Price:price, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}

func (client *NicehashClient) OrderSetPriceDecrease(algo AlgoType, location Location, order uint) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.set.price.decrease", Algo:algo, Location:location, Order:order, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}

func (client *NicehashClient) OrderSetLimit(algo AlgoType, location Location, order uint, limit float32) (string, error) {
	stats := &struct {
		Result struct {
			       Success string `json:"success"`
		       } `json:"result"`
	}{}
	params := &Params{Method:"orders.set.price.limit", Algo:algo, Location:location, Order:order, Limit:limit, ApiId:client.apiid, ApiKey:client.apikey}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return stats.Result.Success, err
	}

	return stats.Result.Success, nil
}
