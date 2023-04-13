package main

import (
	"time"

	"github.com/tbdsux/reqgo"
)

func timeFunc(f func() (*reqgo.Response, error)) (*reqgo.Response, int64, error) {
	start := time.Now()

	res, err := f()
	if err != nil {
		return nil, -1, err
	}

	elapsed := time.Since(start)

	return res, elapsed.Milliseconds(), nil
}
