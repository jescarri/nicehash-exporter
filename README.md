Prometheus Nicehash Exporter
=============================

You can now use prometheus to monitor your nicehash mining operations!

[![Build Status](https://travis-ci.org/jescarri/nicehash-exporter.svg?branch=master)](https://travis-ci.org/jescarri/nicehash-exporter)

Usage
=====

Via Docker:

````
docker run -p 8080:8080 \
-e NICEHASH_ADDR=NICE_HASH_ADDRESS \
-e NICEHASH_ALGO_LIST=[COMMA SEPARATED LIST OF NH ALGO NAMES] EX Keccak,Lyra2REv2,equihash,nist5 \
-e NICEHASH_API_ID=API_ID \
-e NICEHASH_RO_API_KEY=READ_ONLY_NICEHASH_KEY \
-e NICEHASH_WORKERS_TO_MONITOR=[COMMA SEPARATED LIST OF WORKER NAMES] \
-e NICEHASH_FIAT_TO_TRACK=[COMMA SEPATATED LIST OF FIAT CURRENCY] ex MXN,USD \
-e NICEHASH_REFRESH_SECONDS=60 \
jescarri/nicehash-exporter:v1.0.0
````

From Source:

````
export NICEHASH_ADDR=NICE_HASH_ADDRESS
export NICEHASH_ALGO_LIST=[COMMA SEPARATED LIST OF NH ALGO NAMES] EX Keccak,Lyra2REv2,equihash,nist5
export NICEHASH_API_ID=API_ID
export NICEHASH_RO_API_KEY=READ_ONLY_NICEHASH_KEY
export NICEHASH_WORKERS_TO_MONITOR=[COMMA SEPARATED LIST OF WORKER NAMES]
export NICEHASH_FIAT_TO_TRACK=[COMMA SEPATATED LIST OF FIAT CURRENCY] ex MXN,USD
export NICEHASH_REFRESH_SECONDS=60
go run main.go

````

Output
======

````
curl -s localhost:8080/metrics | grep nice
# HELP nicehash_api_call_count Number of api calls made to nicehash and it's status
# TYPE nicehash_api_call_count counter
nicehash_api_call_count{method="GetAlgoStats",status="ok"} 1
nicehash_api_call_count{method="GetFiatBalance",status="ok"} 1
# HELP nicehash_balance Unpaid, paid and unconfirmed balance in BTC and specified FIAT
# TYPE nicehash_balance gauge
nicehash_balance{coin="BTC",status="confirmed"} 0.01196059
nicehash_balance{coin="BTC",status="pending"} 0
nicehash_balance{coin="BTC",status="unpaid"} 0.00042228000000000003
nicehash_balance{coin="MXN",status="confirmed"} 1799.0347493679042
nicehash_balance{coin="MXN",status="pending"} 0
nicehash_balance{coin="MXN",status="unpaid"} 63.516632035968
nicehash_balance{coin="USD",status="confirmed"} 96.327631613975
nicehash_balance{coin="USD",status="pending"} 0
nicehash_balance{coin="USD",status="unpaid"} 3.4009386057
# HELP nicehash_worker_status Worker status 1 is online 0 is offline
# TYPE nicehash_worker_status gauge
nicehash_worker_status{worker_name="R1"} 0
nicehash_worker_status{worker_name="rig1"} 0
nicehash_worker_status{worker_name="rig2"} 0
````

Donations
========

BTC: `18fcigYwqmWBihm9c88zeQG2YxytchsurP`

ETH: `0x128FFde4F0E6A1987d23F424dfc887eebE2d5cC2`

LTC: `LQwwqATDw5iZrqiQ2jFG1SCUKt7SKaAoRs`
