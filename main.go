package main

import (
	"github.com/jescarri/nicehash-exporter/pkg/nhexporter"
	"os"
	"strconv"
)

func main() {
	exp := &nhexporter.NiceHashExporter{}
	addr := os.Getenv("NICEHASH_ADDR")
	algoList := os.Getenv("NICEHASH_ALGO_LIST")
	apiID := os.Getenv("NICEHASH_API_ID")
	roApiKey := os.Getenv("NICEHASH_RO_API_KEY")
	workersToMonitor := os.Getenv("NICEHASH_WORKERS_TO_MONITOR")
	fiatToTrack := os.Getenv("NICEHASH_FIAT_TO_TRACK")
	rs := os.Getenv("NICEHASH_REFRESH_SECONDS")
	refreshSeconds, err := strconv.Atoi(rs)
	if err != nil {
		panic(err)
	}
	exp, err = nhexporter.NewExporter(addr, algoList, apiID, roApiKey, workersToMonitor, fiatToTrack, int64(refreshSeconds))
	if err != nil {
		panic(err)
	}
	exp.CollectMetrics()
}
