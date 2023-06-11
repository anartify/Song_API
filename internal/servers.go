package internal

import (
	"Song_API/internal/routes"
	"Song_API/pkg/cache"
	"Song_API/pkg/controllers"
	"Song_API/pkg/database"
	"Song_API/pkg/middleware"
	"Song_API/pkg/repository"

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
	songCache := cache.NewCacheClient(database.SongCache())
	accountCache := cache.NewCacheClient(database.AccountCache())
	handler := controllers.NewController(repository.SongRepo{}, repository.AccountRepo{}, songCache, accountCache)
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
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
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
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/:id",
		Group:       "songs",
		Version:     "v1",
		Method:      "DELETE",
		Handler:     handler.DeleteSong,
		Middlewares: []gin.HandlerFunc{middleware.Authorization()},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/new",
		Group:       "accounts",
		Version:     "v1",
		Method:      "POST",
		Handler:     handler.CreateAccount,
		Middlewares: []gin.HandlerFunc{},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:        "/",
		Group:       "accounts",
		Version:     "v1",
		Method:      "POST",
		Handler:     handler.GetAccount,
		Middlewares: []gin.HandlerFunc{},
	})
	routes.InitializeRoutes(s.router)
	return s.router.Run()
}
