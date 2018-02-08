# Coindesk

![CoinDesk](https://img.shields.io/badge/docs-GoDoc-ff69b4.svg?style=flat-square)
[![Travis Branch](https://img.shields.io/travis/gostudent/coindesk.svg?style=flat-square)](https://travis-ci.org/gostudent.coindesk)
[![OpenCode](https://img.shields.io/badge/Open-Code-ff6a00.svg?style=flat-square)](https://opencode18.github.io)

coindesk is a Go package to access Coindesk API

## Documentation:
[GoDoc](https://godoc.org/github.com/gostudent/coindesk)

## Installations & Usage

### Installation

`go get github.com/gostudent/coindesk`

### Usage

```go
package main

import (
	"fmt"
	"github.com/gostudent/coindesk"
)

func main() {
	priceUsd := coindesk.GetPrice()
	fmt.Println(priceUsd)

	priceInr := coindesk.GetPrice("INR")
	fmt.Println(priceInr)

	USD, GBP, EUR := coindesk.CurrentPrice()
	fmt.Println(USD, GBP, EUR)
}
```
