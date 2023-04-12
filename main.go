package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	c.AddFunc("@hourly", func() {

	})

	c.Start()
	fmt.Scanln()
}
