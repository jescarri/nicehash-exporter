package nhexporter

import (
	"github.com/bitbandi/go-nicehash-api"
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "nicehash"

type NiceHashExporter struct {
	Addr             string
	Client           *nicehash.NicehashClient
	ApiID            string
	ApiKey           string
	WorkersToMonitor []string
	FiatToTrack      []string
	CollectTime      int64
	AlgosToCheck     []nicehash.AlgoType
}

var (
	niceashApiCallErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "api_call_count",
		Help:      "Number of api calls made to nicehash and it's status",
	},
		[]string{"method", "status"},
	)
)

var (
	nicehashBalance = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "balance",
		Help:      "Unpaid, paid and unconfirmed balance in BTC and specified FIAT",
	},
		[]string{"coin", "status"},
	)
)

var (
	nicehashSpeed = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "speed",
		Help:      "Nicehash reported speed by worker, algo and status, it can be accepted or rejected",
	},
		[]string{"worker_name", "algo", "status"},
	)
)

var (
	nicehashDificulty = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "dificulty",
		Help:      "dificylty by algo",
	},
		[]string{"algo", "worker_name"},
	)
)

var (
	nicehashWorkerStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "worker_status",
		Help:      "Worker status 1 is online 0 is offline",
	},
		[]string{"worker_name"},
	)
)
