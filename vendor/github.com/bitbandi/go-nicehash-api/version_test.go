package nicehash

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	sampleItem := `{"result":{"api_version":"1.0.1"},"method":null}`

	expectedItem := "1.0.1"

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sampleItem)
	})

	nicehashClient := NewNicehashClient(httpClient, "", "FAKEID", "FAKEKEY", "")
	version, err := nicehashClient.GetVersion()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, version)
}
