package internal

import (
	"Song_API/internal/routes"
	"Song_API/pkg/cache"
	"Song_API/pkg/controllers"
	"Song_API/pkg/database"
	"Song_API/pkg/middleware"
	"Song_API/pkg/ratelimit"
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
	bucketCache := cache.NewCacheClient(database.BucketCache())
	handler := controllers.NewController(repository.SongRepo{}, repository.AccountRepo{}, songCache, accountCache)
	globalRule := ratelimit.Rule{Capacity: 1000, Rate: 500}
	rateRules := []ratelimit.Rule{
		{Capacity: 100, Rate: 90, Path: "/v1/songs/", Method: "GET"},
		{Capacity: 50, Rate: 30, Path: "/v1/songs/", Method: "POST"},
		{Capacity: 100, Rate: 90, Path: "/v1/songs/:id", Method: "GET"},
		{Capacity: 80, Rate: 50, Path: "/v1/songs/:id", Method: "PUT"},
		{Capacity: 100, Rate: 90, Path: "/v1/songs/:id", Method: "DELETE"},
		{Capacity: 50, Rate: 30, Path: "/v1/accounts/new", Method: "POST"},
		{Capacity: 100, Rate: 90, Path: "/v1/accounts/", Method: "POST"},
		{Capacity: 100, Rate: 90, Path: "/v1/accounts/all", Method: "GET"},
		{Capacity: 80, Rate: 60, Path: "/v1/accounts/role", Method: "PUT"},
		{Capacity: 50, Rate: 30, Path: "/v1/accounts/", Method: "DELETE"},
	}
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/",
		Group:   "songs",
		Version: "v1",
		Method:  "GET",
		Handler: handler.GetAllSong,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/",
		Group:   "songs",
		Version: "v1",
		Method:  "POST",
		Handler: handler.AddSong,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/:id",
		Group:   "songs",
		Version: "v1",
		Method:  "GET",
		Handler: handler.GetSongById,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/:id",
		Group:   "songs",
		Version: "v1",
		Method:  "PUT",
		Handler: handler.UpdateSong,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/:id",
		Group:   "songs",
		Version: "v1",
		Method:  "DELETE",
		Handler: handler.DeleteSong,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/new",
		Group:   "accounts",
		Version: "v1",
		Method:  "POST",
		Handler: handler.CreateAccount,
		Middlewares: []gin.HandlerFunc{
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/",
		Group:   "accounts",
		Version: "v1",
		Method:  "POST",
		Handler: handler.GetAccount,
		Middlewares: []gin.HandlerFunc{
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/all",
		Group:   "accounts",
		Version: "v1",
		Method:  "GET",
		Handler: handler.GetAllAccount,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/role",
		Group:   "accounts",
		Version: "v1",
		Method:  "PUT",
		Handler: handler.UpdateRole,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/",
		Group:   "accounts",
		Version: "v1",
		Method:  "DELETE",
		Handler: handler.DeleteAccount,
		Middlewares: []gin.HandlerFunc{
			middleware.Authorization([]string{"admin", "general"}, accountCache),
			middleware.RateLimit(rateRules, globalRule, bucketCache),
		},
	})
	routes.InitializeRoutes(s.router)
	return s.router.Run()
}
