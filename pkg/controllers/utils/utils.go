package utils

import (
	"context"
)

// AppReq is a struct to hold the request body, headers, query and params.
type AppReq struct {
	Body    map[string]interface{}
	Headers map[string]string
	Query   map[string]string
	Params  map[string]string
}

type AppResp map[string]interface{}

type RouteHandler func(ctx context.Context, req *AppReq) AppResp
