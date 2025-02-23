package common

import "net/http"

type ApiRequest interface {
	// parse is used to fill request from URL where needed
	Parse(*http.Request) error
	// getNewObj returns new empty obj
	GetNewObj() ApiRequest
	// validate fields
	Validate() error
}

type ApiResponse struct {
	Resp Response
	Code int
	Err  error
}

type Response interface{}
