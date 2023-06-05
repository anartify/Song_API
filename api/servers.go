package api

import (
	"Song_API/api/controllers"
	"Song_API/api/middleware"
	"Song_API/api/repository"
	"Song_API/api/routes"

	"github.com/gin-gonic/gin"
)

// Server struct holds a router field of type *gin.Engine.
type Server struct {
	router *gin.Engine
}

// NewServer() instantiates a new Server object with router value as gin.Default() and returns a pointer to it
func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

// Start() registers the routes and starts the server
func (s *Server) Start() error {
	handler := &controllers.Controller{SongRepo: repository.SongRepo{}, AccountRepo: repository.AccountRepo{}}
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "songs",
		Version:     "v1",
		Method:      "GET",
		Handler:     handler.GetAllSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "songs",
		Version:     "v1",
		Method:      "POST",
		Handler:     handler.AddSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization(), middleware.Validation()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "GET",
		Handler:     handler.GetSongById,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "PUT",
		Handler:     handler.UpdateSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization(), middleware.Validation()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "DELETE",
		Handler:     handler.DeleteSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization(), middleware.Validation()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/new",
		Group:       "accounts",
		Version:     "v1",
		Method:      "POST",
		Handler:     handler.CreateAccount,
		Middlewares: []gin.HandlerFunc{middleware.Validation()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "accounts",
		Version:     "v1",
		Method:      "POST",
		Handler:     handler.GetAccount,
		Middlewares: []gin.HandlerFunc{middleware.Validation()},
	})
	routes.InitializeRoutes(s.router)
	return s.router.Run()
}
