package main

import (
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

var (
	d, _ = deta.New()
)

func getDB(dbName string) *base.Base {
	db, _ := base.New(d, dbName)
	return db
}

func handleFetcher(eUrl string, typ string, fetcher func(string, string) (FetchData, error)) error {
	f, err := fetcher(eUrl, typ)
	if err != nil {
		return err
	}

	return saveData(f, typ)
}

// save data to custom base
// dbName is report type [account_http, account_https, api_http,...]
func saveData(f FetchData, dbName string) error {
	db := getDB(dbName)

	_, err := db.Put(&f)
	return err
}
