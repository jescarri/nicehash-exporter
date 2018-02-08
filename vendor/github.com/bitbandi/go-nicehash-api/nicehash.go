package nicehash

import (
	"github.com/dghubble/sling"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"log"
	"strings"
)

type NicehashClient struct {
	sling      *sling.Sling
	apiid      string
	apikey     string
	httpClient *nicehashHttpClient
}

// server send the api response with text/html content type
// we fix this: change content type to json
type nicehashHttpClient struct {
	client    *http.Client
	debug     bool
	useragent string
}

type Params struct {
	Method   string `url:"method"`
	ApiId    string `url:"id,omitempty"`
	ApiKey   string `url:"key,omitempty"`
	Addr     string `url:"addr,omitempty"`
	Algo     AlgoType `url:"algo"`
	Location Location `url:"location"`
	My       bool `url:"my,omitempty"`

	Order    uint `url:"order,omitempty"`
	Limit    float32 `url:"limit,omitempty"`
	Price    float32 `url:"price,omitempty"`
	Amount   float64 `url:"amount,omitempty"`
}

func (d nicehashHttpClient) Do(req *http.Request) (*http.Response, error) {
	if d.debug {
		d.dumpRequest(req)
	}
	if d.useragent != "" {
		req.Header.Set("User-Agent", d.useragent)
	}
	client := func() (*http.Client) {
		if d.client != nil {
			return d.client
		} else {
			return http.DefaultClient
		}
	}()
	if client.Transport != nil {
		if transport, ok := client.Transport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	} else {
		if transport, ok := http.DefaultTransport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	}
	resp, err := client.Do(req)
	if d.debug {
		d.dumpResponse(resp)
	}
	if err == nil {
		contenttype := resp.Header.Get("Content-Type");
		if len(contenttype) == 0 || strings.HasPrefix(contenttype, "text/html") {
			resp.Header.Set("Content-Type", "application/json")
		}
	}
	return resp, err
}

func (d nicehashHttpClient) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (d nicehashHttpClient) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func NewNicehashClient(client *http.Client, BaseURL string, ApiId string, ApiKey string, UserAgent string) *NicehashClient {
	if len(BaseURL) == 0 {
		BaseURL = "https://api.nicehash.com/"
	}
	nicehashclient := &nicehashHttpClient{client:client, useragent:UserAgent}
	return &NicehashClient{
		httpClient: nicehashclient,
		sling: sling.New().Doer(nicehashclient).Base(strings.TrimRight(BaseURL, "/") + "/").Path("api"),
		apiid: ApiId,
		apikey: ApiKey,
	}
}

func (client NicehashClient) SetDebug(debug bool) {
	client.httpClient.debug = debug
}
