package services

import (
	"log"
	"net/http"
)

type FacebookClient interface {
	GetRequest(url string) (FbResponse, error)
}

type FbClient struct {
}

type FbResponse struct {
	Code int
	Body string
}

func (c *FbClient) GetRequest(url string) (FbResponse, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Http request error: %s", err.Error())
	}
	defer resp.Body.Close()

	return FbResponse{Code: resp.StatusCode, Body: HttpResponseBodyToString(resp.Body)}, err
}
