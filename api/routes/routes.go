package routes

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AppReq struct {
	Body    map[string]interface{}
	Headers map[string]string
	Query   map[string]string
	Params  map[string]string
}

type AppResp map[string]interface{}

type RouteHandler func(ctx context.Context, req *AppReq) AppResp

type RouteDef struct {
	Path    string
	Version string
	Method  string
	Handler RouteHandler
}

var clientRoutes = []RouteDef{}

func RegisterRoutes(r RouteDef) {
	clientRoutes = append(clientRoutes, r)
}

func (r *RouteDef) GetPath() string {
	return "/" + r.Version + r.Path
}

func InitializeRoutes(server *gin.Engine) {
	for _, r := range clientRoutes {
		route := r
		server.Handle(route.Method, route.GetPath(), func(c *gin.Context) {
			appReq := &AppReq{
				Body:    make(map[string]interface{}),
				Headers: make(map[string]string),
				Query:   make(map[string]string),
				Params:  make(map[string]string),
			}

			for key, values := range c.Request.Header {
				if len(values) > 0 {
					appReq.Headers[key] = values[0]
				}
			}

			for key, values := range c.Request.URL.Query() {
				if len(values) > 0 {
					appReq.Query[key] = values[0]
				}
			}

			for _, param := range c.Params {
				appReq.Params[param.Key] = param.Value
			}
			if c.Request.ContentLength > 0 {
				if err := c.BindJSON(&appReq.Body); err != nil {
					panic("failed to bind json")
				}
			}
			resp := route.Handler(c.Request.Context(), appReq)
			c.JSON(resp["status"].(int), resp)
		})
	}
}
