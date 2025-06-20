package domain

import "backend-api/client/http"

var d *Domain

type Domain struct {
	WsHttpClient *http.WSClient
}

func NewDomain(wsHttpClient *http.WSClient) {
	d = &Domain{
		WsHttpClient: wsHttpClient,
	}
}
