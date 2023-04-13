package main

import (
	"encoding/json"
	"os"
)

type Endpoints struct {
	Meta   Meta            `json:"meta"`
	Report EndpointsReport `json:"report"`
}

type ArrReportElements = [][]interface{}

type EndpointsReport struct {
	AccountHTTP             ArrReportElements `json:"account_http"`
	AccountHTTPS            ArrReportElements `json:"account_https"`
	ApiHTTP                 ArrReportElements `json:"api_http"`
	ApiHTTPS                ArrReportElements `json:"api_https"`
	ApiHTTPS2               ArrReportElements `json:"api_https2"`
	AtomicHTTP              ArrReportElements `json:"atomic_http"`
	AtomicHTTPS             ArrReportElements `json:"atomic_https"`
	DFuseHTTPS              ArrReportElements `json:"dfuse_https"`
	FirehoseHTTPS           ArrReportElements `json:"firehose_https"`
	HistoryTraditionalHTTP  ArrReportElements `json:"history_traditional_http"`
	HistoryTraditionalHTTPS ArrReportElements `json:"history_traditional_https"`
	HyperionHTTP            ArrReportElements `json:"hyperion_http"`
	HyperionHTTPS           ArrReportElements `json:"hyperion_https"`
	P2P                     ArrReportElements `json:"p2p"`
}

type Meta struct {
	Details Details `json:"details"`
	Network Network `json:"network"`
	Title   Details `json:"title"`
	Update  Network `json:"update"`
}

type Details struct {
	Value string `json:"value"`
}

type Network struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ReportClass struct {
	HTMLName string `json:"html_name"`
	Name     string `json:"name"`
	Rank     int64  `json:"rank"`
}

// parse `enpoints.json` in directory
func parseEndpoints() (Endpoints, error) {
	var data Endpoints

	file, err := os.ReadFile("endpoints.json")
	if err != nil {
		return Endpoints{}, err
	}

	if err = json.Unmarshal(file, &data); err != nil {
		return Endpoints{}, err
	}

	return data, nil
}
