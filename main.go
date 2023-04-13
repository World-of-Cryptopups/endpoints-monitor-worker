package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("Starting...")

	// parse `endpoints.json`
	log.Println("Parsing endpoints...")
	endpoints, err := parseEndpoints()
	if err != nil {
		log.Fatalln("Failed to parse endpoints.json. Quitting...")
	}
	log.Println("Successfully parsed endpoints.")

	c := cron.New()

	// account_http
	c.AddFunc("@hourly", func() {
		log.Println("account_http")
		for _, v := range endpoints.Report.AccountHTTP {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "account_http", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// account_https
	c.AddFunc("@hourly", func() {
		log.Println("account_https")
		for _, v := range endpoints.Report.AccountHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "account_https", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// api_http
	c.AddFunc("@hourly", func() {
		log.Println("api_http")
		for _, v := range endpoints.Report.ApiHTTP {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "api_http", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// api_https
	c.AddFunc("@hourly", func() {
		log.Println("api_https")
		for _, v := range endpoints.Report.ApiHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "api_https", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// api_https2
	c.AddFunc("@hourly", func() {
		log.Println("api_https2")
		for _, v := range endpoints.Report.ApiHTTPS2 {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "api_https2", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// atomic_http
	c.AddFunc("@hourly", func() {
		log.Println("atomic_http")
		for _, v := range endpoints.Report.AtomicHTTP {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "atomic_http", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// atomic_https
	c.AddFunc("@hourly", func() {
		log.Println("atomic_https")
		for _, v := range endpoints.Report.AtomicHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "atomic_https", fetchAtomicEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// dfuse_https
	c.AddFunc("@hourly", func() {
		log.Println("dfuse_https")
		for _, v := range endpoints.Report.DFuseHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "dfuse_https", fetchGet); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// firehose_https
	c.AddFunc("@hourly", func() {
		log.Println("firehose_https")
		for _, v := range endpoints.Report.FirehoseHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "firehose_https", fetchGet); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// history_traditional_http
	c.AddFunc("@hourly", func() {
		log.Println("history_traditional_http")
		for _, v := range endpoints.Report.HistoryTraditionalHTTP {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "history_traditional_http", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// history_traditional_https
	c.AddFunc("@hourly", func() {
		log.Println("history_traditional_https")
		for _, v := range endpoints.Report.HistoryTraditionalHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "history_traditional_https", fetchApiEndpoint); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// hyperion_http
	c.AddFunc("@hourly", func() {
		log.Println("hyperion_http")
		for _, v := range endpoints.Report.HyperionHTTP {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "hyperion_http", fetchHyperion); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// hyperion_https
	c.AddFunc("@hourly", func() {
		log.Println("hyperion_https")
		for _, v := range endpoints.Report.HyperionHTTPS {
			eUrl := v[1].(string)

			if err := handleFetcher(eUrl, "hyperion_https", fetchHyperion); err != nil {
				log.Println(err)
				continue
			}
		}
	})

	// p2p
	// still not sure for testing implemention of p2p
	// TODO:

	c.Run()
}
