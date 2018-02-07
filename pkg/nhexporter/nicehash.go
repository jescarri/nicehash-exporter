package nhexporter

import (
	"fmt"
	"github.com/bitbandi/go-nicehash-api"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/gostudent/coindesk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewExporter(niceHasAddr string, algosToCheck string, apiId string, apiKey string, rigsToMonitor string, fiatToTrack string, collectTime int64) (*NiceHashExporter, error) {
	exp := &NiceHashExporter{}
	exp.Client = nicehash.NewNicehashClient(&http.Client{}, "", apiId, apiKey, "nicehash-exporter")
	exp.Addr = niceHasAddr
	exp.ApiID = apiId
	exp.ApiKey = apiKey
	exp.WorkersToMonitor = strings.Split(rigsToMonitor, ",")
	if collectTime == 0 {
		exp.CollectTime = 30
	} else {
		exp.CollectTime = collectTime
	}
	if len(exp.WorkersToMonitor) == 0 {
		return &NiceHashExporter{}, fmt.Errorf("Not enough rigs to monitor")
	}
	algos := strings.Split(algosToCheck, ",")
	if len(algosToCheck) == 0 {
		return &NiceHashExporter{}, fmt.Errorf("Not enough algoritms")
	}
	for _, algo := range algos {
		algoId, err := AlgoFromStringToInt(algo)
		if err != nil {
			return &NiceHashExporter{}, fmt.Errorf("%s is not a valid NiceHash algoritm", algo)
		}
		exp.AlgosToCheck = append(exp.AlgosToCheck, algoId...)
	}
	exp.FiatToTrack = strings.Split(fiatToTrack, ",")
	if len(exp.FiatToTrack) == 0 {
		exp.FiatToTrack = []string{"USD"}
	}
	prometheus.MustRegister(niceashApiCallErrors)
	prometheus.MustRegister(nicehashBalance)
	prometheus.MustRegister(nicehashSpeed)
	prometheus.MustRegister(nicehashWorkerStatus)
	return exp, nil
}

func (e *NiceHashExporter) GetAlgoStats() (map[string][]nicehash.ProviderWorker, error) {
	workerStats := map[string][]nicehash.ProviderWorker{}
	for _, algo := range e.AlgosToCheck {
		wrkrStat, err := e.Client.GetStatsProviderWorkers(e.Addr, algo)
		if err != nil {
			return map[string][]nicehash.ProviderWorker{}, err
		}
		workerStats[algo.ToString()] = wrkrStat
	}
	return workerStats, nil
}

func (e *NiceHashExporter) GetFiatBalance() (map[string]map[string]float64, error) {
	balances := make(map[string]map[string]float64)
	balances["confirmed"] = make(map[string]float64)
	balances["pending"] = make(map[string]float64)
	balances["unpaid"] = make(map[string]float64)
	balance, err := e.Client.GetBalance()
	if err != nil {
		return map[string]map[string]float64{}, err
	}
	balances["confirmed"]["BTC"] = balance.Confirmed
	balances["pending"]["BTC"] = balance.Pending

	for _, fiat := range e.FiatToTrack {
		fiat_price := coindesk.GetPrice(strings.ToUpper(fiat))
		balances["confirmed"][strings.ToUpper(fiat)] = balance.Confirmed * fiat_price
		balances["pending"][strings.ToUpper(fiat)] = balance.Pending * fiat_price
	}
	stats, err := e.Client.GetStatsProviderEx(e.Addr)
	if err != nil {
		return map[string]map[string]float64{}, err
	}
	for _, algoStats := range stats.Current {
		balances["unpaid"]["BTC"] += algoStats.Unpaid
		for _, fiat := range e.FiatToTrack {
			fiat_price := coindesk.GetPrice(strings.ToUpper(fiat))
			balances["unpaid"][strings.ToUpper(fiat)] += algoStats.Unpaid * fiat_price
		}
	}
	return balances, nil
}

func (e *NiceHashExporter) DetermineWorkerConnection(stats map[string][]nicehash.ProviderWorker) map[string]int {
	rigsConnStats := map[string]int{}
	for _, rigName := range e.WorkersToMonitor {
		rigsConnStats[rigName] = 0
		//iterate over all Algos, set rig status to connected 1 if rig name present
		for _, algoStats := range stats {
			// iterate over all rigs for a particular algo
			for _, stat := range algoStats {
				if stat.Name == rigName {
					rigsConnStats[rigName] = 1
					continue
				}
			}
		}
	}
	return rigsConnStats
}

func (e *NiceHashExporter) CollectMetrics() {
	go func() {
		// The Handler function provides a default handler to expose metrics
		// via an HTTP server. "/metrics" is the usual endpoint for that.
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
	for {
		balances, err := e.GetFiatBalance()

		if err != nil {
			niceashApiCallErrors.WithLabelValues("GetFiatBalance", "error").Inc()
		}
		niceashApiCallErrors.WithLabelValues("GetFiatBalance", "ok").Inc()
		for status, b := range balances {
			for fiat, val := range b {
				nicehashBalance.WithLabelValues(fiat, status).Set(val)
			}
		}
		as, err := e.GetAlgoStats()
		if err != nil {
			niceashApiCallErrors.WithLabelValues("GetAlgoStats", "error").Inc()
		}
		niceashApiCallErrors.WithLabelValues("GetAlgoStats", "ok").Inc()
		for algo, workers := range as {
			for _, w := range workers {
				nicehashSpeed.WithLabelValues(w.Name, algo, "accepted").Set(w.AcceptedSpeed)
				nicehashSpeed.WithLabelValues(w.Name, algo, "rejected").Set(w.RejectedSpeed)
				nicehashDificulty.WithLabelValues(algo, w.Name).Set(w.Difficulty)

			}
		}
		connStat := e.DetermineWorkerConnection(as)
		for w, s := range connStat {
			nicehashWorkerStatus.WithLabelValues(w).Set(float64(s))
		}
		time.Sleep(time.Duration(e.CollectTime) * time.Second)
	}
	select {}
}
