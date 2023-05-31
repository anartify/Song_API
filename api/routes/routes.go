package routes

import (
	"context"

	"github.com/gin-gonic/gin"
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

// RouteDef struct holds the path, version, request method, middlewares and associated handler function of a route.
type RouteDef struct {
	Path        string
	Version     string
	Method      string
	Handler     RouteHandler
	Middlewares []gin.HandlerFunc
}

var clientRoutes = []RouteDef{}

// RegisterRoutes appends the RouteDef to the clientRoutes slice.
func RegisterRoutes(r RouteDef) {
	clientRoutes = append(clientRoutes, r)
}

// GetPath returns the path of the route.
func (r *RouteDef) GetPath() string {
	return "/" + r.Version + r.Path
}

// InitializeRoutes registers request Handle and middleware for the different clientRoutes. It extracts AppReq from gin.Context and calls the correspoding RoutesHandler function.
func InitializeRoutes(server *gin.Engine) {
	for _, r := range clientRoutes {
		route := r
		ginHandler := func(c *gin.Context) {
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
		}
		handlers := append(route.Middlewares, ginHandler)
		server.Handle(route.Method, route.GetPath(), handlers...)
	}
}
