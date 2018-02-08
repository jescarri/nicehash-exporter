package nicehash

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"balance_confirmed":"0.00500000","balance_pending":"0.00000000"},"method":"balance"}`

	expectedItem := Balance{
		Confirmed: 0.005,
		Pending: 0.00,
	}

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "balance", r.URL.Query().Get("method"))
		assert.Equal(t, "FAKEID", r.URL.Query().Get("id"))
		assert.Equal(t, "FAKEKEY", r.URL.Query().Get("key"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.GetBalance()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}

