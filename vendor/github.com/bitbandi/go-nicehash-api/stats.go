package nicehash

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type GlobalStats struct {
	Algo                  AlgoType `json:"algo"`
	ProfitabilityAboveBtc float32  `json:"profitability_above_btc,string"`
	ProfitabilityAboveLtc float32  `json:"profitability_above_ltc,string"`
	Price                 float64  `json:"price,string"`
	ProfitabilityBtc      float32  `json:"profitability_btc,string"`
	ProfitabilityLtc      float32  `json:"profitability_ltc,string"`
	Speed                 float64  `json:"speed,string"`
}

func (client *NicehashClient) GetStatsGlobalCurrent() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			Error string        `json:"error"`
			Stats []GlobalStats `json:"stats"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.global.current", Algo: AlgoTypeMAX, Location: LocationMAX}
	resp, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, err
	}
	if code := resp.StatusCode; code < 200 || 299 < code {
		return nil, errors.New("Http response: " + resp.Status)
	}
	if stats.Result.Error != "" {
		return nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, nil
}

func (client *NicehashClient) GetStatsGlobalDay() ([]GlobalStats, error) {
	stats := &struct {
		Result struct {
			Error string        `json:"error"`
			Stats []GlobalStats `json:"stats"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.global.24h", Algo: AlgoTypeMAX, Location: LocationMAX}
	resp, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, err
	}
	if code := resp.StatusCode; code < 200 || 299 < code {
		return nil, errors.New("Http response: " + resp.Status)
	}
	if stats.Result.Error != "" {
		return nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, nil
}

type ProviderStats struct {
	Algo          AlgoType `json:"algo"`
	Balance       float64  `json:"balance,string"`
	AcceptedSpeed float64  `json:"accepted_speed,string"`
	RejectedSpeed float64  `json:"rejected_speed,string"`
}

type ProviderPayments struct {
	Amount float64   `json:"amount,string"`
	Fee    float64   `json:"fee,string"`
	TxID   string    `json:"TXID"`
	Time   time.Time `json:"time"`
}

func (t *ProviderPayments) UnmarshalJSON(data []byte) error {
	var err error
	type Alias ProviderPayments
	aux := &struct {
		Time string `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Time, err = time.Parse("2006-01-02 15:04:05", aux.Time)
	return err
}

func (client *NicehashClient) GetStatsProvider(addr string) ([]ProviderStats, []ProviderPayments, error) {
	stats := &struct {
		Result struct {
			Error    string             `json:"error"`
			Stats    []ProviderStats    `json:"stats"`
			Payments []ProviderPayments `json:"payments"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.provider", Algo: AlgoTypeMAX, Location: LocationMAX, Addr: addr}
	resp, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, nil, err
	}
	if code := resp.StatusCode; code < 200 || 299 < code {
		return nil, nil, errors.New("Http response: " + resp.Status)
	}
	if stats.Result.Error != "" {
		return nil, nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Stats, stats.Result.Payments, nil
}

type ProviderExStats struct {
	Algo          AlgoType `json:"algo"`
	Suffix        string   `json:"suffix"`
	Name          string   `json:"name"`
	Profitability float64  `json:"profitability,string"`
	Unpaid        float64  `json:"balance,string"`
	AcceptedSpeed float64  `json:"accepted_speed,string"`
	RejectedSpeed float64  `json:"rejected_speed,string"`
}

func (t *ProviderExStats) UnmarshalJSON(data []byte) error {
	var err error
	type Alias ProviderExStats
	aux := &struct {
		Data []interface{} `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	speeds := aux.Data[0].(map[string]interface{})
	if val, ok := speeds["a"]; ok {
		t.AcceptedSpeed, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return err
		}
	}
	if val, ok := speeds["rs"]; ok {
		t.RejectedSpeed, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return err
		}
	}
	t.Unpaid, err = strconv.ParseFloat(aux.Data[1].(string), 64)
	return err
}

type ProviderExHistory struct {
	Algo AlgoType                            `json:"algo"`
	Data map[time.Time]ProviderExHistoryItem `json:"data"`
}

type ProviderExHistoryItem struct {
	Unpaid        float64 `json:"balance,string"`
	AcceptedSpeed float64 `json:"accepted_speed,string"`
	RejectedSpeed float64 `json:"rejected_speed,string"`
}

func (t *ProviderExHistory) UnmarshalJSON(data []byte) error {
	var err error
	type Alias ProviderExHistory
	aux := &struct {
		Data [][]interface{} `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Data = make(map[time.Time]ProviderExHistoryItem)
	for _, d := range aux.Data {
		var item ProviderExHistoryItem
		date := time.Unix(int64(d[0].(float64))*300, 0)
		speeds := d[1].(map[string]interface{})
		if val, ok := speeds["a"]; ok {
			item.AcceptedSpeed, err = strconv.ParseFloat(val.(string), 64)
			if err != nil {
				return err
			}
		}
		if val, ok := speeds["rs"]; ok {
			item.RejectedSpeed, err = strconv.ParseFloat(val.(string), 64)
			if err != nil {
				return err
			}
		}
		item.Unpaid, err = strconv.ParseFloat(d[2].(string), 64)
		if err != nil {
			return err
		}
		if item.AcceptedSpeed > 0 || item.RejectedSpeed > 0 {
			t.Data[date] = item
		}
	}
	return nil
}

type ProviderExPayments struct {
	Amount float64   `json:"amount,string"`
	Fee    float64   `json:"fee,string"`
	TxID   string    `json:"TXID"`
	Time   time.Time `json:"time"`
}

func (t *ProviderExPayments) UnmarshalJSON(data []byte) error {
	var err error
	type Alias ProviderPayments
	aux := &struct {
		Time int64 `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Time = time.Unix(aux.Time, 0)
	return err
}

type StatsProviderEx struct {
	Error    string               `json:"error"`
	Current  []ProviderExStats    `json:"current"`
	Past     []ProviderExHistory  `json:"past"`
	Payments []ProviderExPayments `json:"payments"`
}

func (client *NicehashClient) GetStatsProviderEx(addr string) (StatsProviderEx, error) {
	stats := &struct {
		Result StatsProviderEx `json:"result"`
	}{}
	params := &Params{Method: "stats.provider.ex", Algo: AlgoTypeMAX, Location: LocationMAX, Addr: addr}
	resp, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return StatsProviderEx{}, err
	}
	if code := resp.StatusCode; code < 200 || 299 < code {
		return StatsProviderEx{}, errors.New("Http response: " + resp.Status)
	}
	if stats.Result.Error != "" {
		return StatsProviderEx{}, errors.New(stats.Result.Error)
	}
	return stats.Result, nil
}

type ProviderWorker struct {
	Name          string   `json:"name"`
	AcceptedSpeed float64  `json:"accepted_speed,string"`
	RejectedSpeed float64  `json:"rejected_speed,string"`
	Connected     uint64   `json:"connected"`
	XnSubEnabled  bool     `json:"xnsub"`
	Difficulty    float64  `json:"difficulty"`
	Location      Location `json:"location"`
}

func (t *ProviderWorker) UnmarshalJSON(data []byte) error {
	var err error
	var aux []interface{}
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t.Name = aux[0].(string)
	speeds := aux[1].(map[string]interface{})
	if val, ok := speeds["a"]; ok {
		t.AcceptedSpeed, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return err
		}
	}
	if val, ok := speeds["rs"]; ok {
		t.RejectedSpeed, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return err
		}
	}
	t.Connected = uint64(aux[2].(float64))
	t.XnSubEnabled = aux[3].(float64) == 1
	t.Difficulty, err = strconv.ParseFloat(aux[4].(string), 64)
	if err != nil {
		return err
	}
	t.Location = Location(uint64(aux[5].(float64)))
	return nil
}

func (client *NicehashClient) GetStatsProviderWorkers(addr string, algo AlgoType) ([]ProviderWorker, error) {
	stats := &struct {
		Result struct {
			Error   string           `json:"error"`
			Address string           `json:"addr"`
			Algo    AlgoType         `json:"algo"`
			Workers []ProviderWorker `json:"workers"`
		} `json:"result"`
	}{}
	params := &Params{Method: "stats.provider.workers", Algo: algo, Location: LocationMAX, Addr: addr}
	_, err := client.sling.New().Get("").QueryStruct(params).ReceiveSuccess(&stats)
	if err != nil {
		return nil, err
	}
	if stats.Result.Error != "" {
		return nil, errors.New(stats.Result.Error)
	}
	return stats.Result.Workers, nil
}
