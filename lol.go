package main

import (
	"github.com/buaazp/stress"
)

func main() {
	config := &stress.Config{
		URL:         "http://88.198.59.166/cc",
		Method:      "GET",
		Concurrency: 10,
		Requests:    1000,
	}

	stress.NewRunner(config).Run()
}
