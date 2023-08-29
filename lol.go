package main

import (
	"github.com/utkusen/reqstress"
)

func main() {
	config := &reqstress.Config{
		URL:         "http://88.198.59.166/cc",
		Method:      "GET",
		Concurrency: 10,
		Requests:    1000,
	}

	reqstress.NewRunner(config).Run()
}