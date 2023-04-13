package main

import (
	"fmt"
	"time"

	"github.com/tbdsux/reqgo"
)

type FetchData struct {
	Url      string            `json:"url"`
	MS       int64             `json:"ms"`
	Response FetchDataResponse `json:"response"`
	IsError  bool              `json:"is_error"`
	Error    string            `json:"error"`
	DateTime int64             `json:"date_time"`
	Type     string            `json:"type"`
}

type FetchDataResponse struct {
	Code   int
	Status string
}

func handleFetch(url, typ string, f func() (*reqgo.Response, error)) (FetchData, error) {
	r, ms, err := timeFunc(f)

	isError := false
	errString := ""
	if err != nil {
		isError = true
		errString = err.Error()
	}

	return FetchData{
			Url: url,
			MS:  ms,
			Response: FetchDataResponse{
				Code:   r.Code,
				Status: r.Status,
			},
			IsError:  isError,
			Error:    errString,
			DateTime: time.Now().Unix(),
			Type:     typ,
		},
		nil
}

func fetchApiEndpoint(url string, typ string) (FetchData, error) {
	finalUrl := fmt.Sprintf("%s/v1/chain/get_info", url)

	return handleFetch(url, typ, func() (*reqgo.Response, error) {
		r, err := reqgo.Post(finalUrl, nil)
		if err != nil {
			return nil, err
		}

		return r, err
	})
}

func fetchAtomicEndpoint(url string, typ string) (FetchData, error) {
	finalUrl := fmt.Sprintf("%s/atomicassets/v1/config", url)

	return handleFetch(url, typ, func() (*reqgo.Response, error) {
		r, err := reqgo.Get(finalUrl, nil)
		if err != nil {
			return nil, err
		}

		return r, err
	})
}

func fetchGet(url, typ string) (FetchData, error) {
	return handleFetch(url, typ, func() (*reqgo.Response, error) {
		r, err := reqgo.Get(url, nil)
		if err != nil {
			return nil, err
		}

		return r, err
	})
}

func fetchHyperion(url, typ string) (FetchData, error) {
	finalUrl := fmt.Sprintf("%s/v2/health", url)

	return handleFetch(url, typ, func() (*reqgo.Response, error) {
		return reqgo.Get(finalUrl, nil)
	})
}
