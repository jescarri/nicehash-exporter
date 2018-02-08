package nicehash

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetOrders(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
	   "result":{
	      "orders":[
		 {
		    "type":0,
		    "id":5877,
		    "price":"0.0505",
		    "algo":1,
		    "alive":true,
		    "limit_speed":"1.0",
		    "workers":0,
		    "accepted_speed":"0.0"
		 }
	      ]
	   },
	   "method":"orders.get"
	}`

	expectedItem := []Orders{
		{
			Type: 0,
			Id: 5877,
			Price: 0.0505,
			Algo: 1,
			Alive: true,
			LimitSpeed: 1.0,
			Workers: 0,
			AcceptedSpeed: 0.0,
		},
	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "orders.get", r.URL.Query().Get("method"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	stats, err := nicehashClient.GetOrders(0, 0)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, stats)
}

func TestGetMyOrders(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{
	   "result":{
	      "orders":[
		 {
		    "type":0,
		    "btc_avail":"0.01751439",
		    "limit_speed":"0.0",
		    "pool_user":"worker",
		    "pool_port":3333,
		    "alive":false,
		    "workers":0,
		    "pool_pass":"x",
		    "accepted_speed":"0.0",
		    "id":1879,
		    "algo":0,
		    "price":"1.0000",
		    "btc_paid":"0.00000000",
		    "pool_host":"testpool.com",
		    "end":1413294447421
		 }
	      ]
	   },
	   "method":"orders.get"
	}`

	expectedItem := []MyOrders{
		{
			Type: 0,
			BtcAvail: 0.01751439,
			LimitSpeed: 0.0,
			PoolUser: "worker",
			PoolPort: 3333,
			Alive: false,
			Workers: 0,
			PoolPass: "x",
			AcceptedSpeed: 0.0,
			Id: 1879,
			Algo: 0,
			Price: 1.0000,
			BtcPaid: 0.00000000,
			PoolHost: "testpool.com",
			End: 1413294447421,
		},
	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "orders.get", r.URL.Query().Get("method"))
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	stats, err := nicehashClient.GetMyOrders(0, 0)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, stats)
}

func TestOrderRefill(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"success":"Order #123 refilled."},"method":"orders.refill"}`

	expectedItem := "Order #123 refilled."

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		assert.Equal(t, "123", r.URL.Query().Get("order"))
		assert.Equal(t, "0.01", r.URL.Query().Get("amount"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.OrderRefill(0, 0, 123, 0.01)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}

func TestOrderRemove(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"success":"Order removed."},"method":"orders.remove"}`

	expectedItem := "Order removed."

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		assert.Equal(t, "123", r.URL.Query().Get("order"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.OrderRemove(0, 0, 123)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}

func TestOrderSetPrice(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"success":"New order price set to: 2.10"},"method":"orders.set.price"}`

	expectedItem := "New order price set to: 2.10"

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		assert.Equal(t, "123", r.URL.Query().Get("order"))
		assert.Equal(t, "2.1", r.URL.Query().Get("price"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.OrderSetPrice(0, 0, 123, 2.1)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}

func TestOrderSetPriceDecrease(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"success":"New order price set to: 2.10"},"method":"orders.set.price"}`

	expectedItem := "New order price set to: 2.10"

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		assert.Equal(t, "123", r.URL.Query().Get("order"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.OrderSetPriceDecrease(0, 0, 123)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}

func TestOrderSetLimit(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"success":"New order limit set to: 1.00"},"method":"orders.set.limit"}`

	expectedItem := "New order limit set to: 1.00"

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		assert.Equal(t, "123", r.URL.Query().Get("order"))
		assert.Equal(t, "1", r.URL.Query().Get("limit"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.OrderSetLimit(0, 0, 123, 1.0)

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}
