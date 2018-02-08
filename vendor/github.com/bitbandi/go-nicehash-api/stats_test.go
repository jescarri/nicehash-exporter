package nicehash

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetStatsGlobalCurrent(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
	   "result":{
	      "stats":[
		 {
		    "profitability_above_ltc":"8.27",
		    "price":"0.1683",
		    "profitability_ltc":"0.1554",
		    "algo":0,
		    "speed":"27.0678"
		 },
		 {
		    "price":"0.0117",
		    "profitability_btc":"0.0114",
		    "profitability_above_btc":"2.39",
		    "algo":1,
		    "speed":"1597723.0669"
		 }
	      ]
	   },
	   "method":"stats.global.current"
	}`

	expectedItem := []GlobalStats{
		{
			ProfitabilityAboveLtc: 8.27,
			Price: 0.1683,
			ProfitabilityLtc: 0.1554,
			Algo:0,
			Speed: 27.0678,
		},
		{
			Price: 0.0117,
			ProfitabilityBtc: 0.0114,
			ProfitabilityAboveBtc: 2.39,
			Algo:1,
			Speed: 1597723.0669,
		},

	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "stats.global.current", r.URL.Query().Get("method"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	stats, err := nicehashClient.GetStatsGlobalCurrent()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, stats)
}

func TestGetStatsGlobalDay(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
	   "result":{
	      "stats":[
		 {
		    "profitability_above_ltc":"8.27",
		    "price":"0.1683",
		    "profitability_ltc":"0.1554",
		    "algo":0,
		    "speed":"27.0678"
		 },
		 {
		    "price":"0.0117",
		    "profitability_btc":"0.0114",
		    "profitability_above_btc":"2.39",
		    "algo":1,
		    "speed":"1597723.0669"
		 }
	      ]
	   },
	   "method":"stats.global.current"
	}`

	expectedItem := []GlobalStats{
		{
			ProfitabilityAboveLtc: 8.27,
			Price: 0.1683,
			ProfitabilityLtc: 0.1554,
			Algo:0,
			Speed: 27.0678,
		},
		{
			Price: 0.0117,
			ProfitabilityBtc: 0.0114,
			ProfitabilityAboveBtc: 2.39,
			Algo:1,
			Speed: 1597723.0669,
		},

	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "stats.global.24h", r.URL.Query().Get("method"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	stats, err := nicehashClient.GetStatsGlobalDay()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, stats)
}