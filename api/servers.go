package api

import (
	"Song_API/api/controllers"
	"Song_API/api/repository"
	"Song_API/api/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	handler := &controllers.Controller{Repo: repository.SongRepo{}}
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs",
		Version: "v1",
		Method:  "GET",
		Handler: handler.GetAllSong,
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs",
		Version: "v1",
		Method:  "POST",
		Handler: handler.AddSong,
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "GET",
		Handler: handler.GetSongById,
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "PUT",
		Handler: handler.UpdateSong,
	})
	routes.RegisterRoutes(routes.RouteDef{
		Path:    "/songs/:id",
		Version: "v1",
		Method:  "DELETE",
		Handler: handler.DeleteSong,
	})
	routes.InitializeRoutes(s.router)
	return s.router.Run()
}
