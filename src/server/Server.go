package server

import (
	database "api-order/src/Database"
	"api-order/src/client/infraestructure/http/routes"

	"log"
	"api-order/src/config" 
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine 		*gin.Engine
	http 		string
	port 		string
	httpAddr 	string
}

func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine: gin.New(),
		http: http,
		port: port,
		httpAddr: http + ":" + port,
	}

	database.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	srv.engine.Use(config.ConfigurationCors())

	return srv
}

func (s *Server) registerRoutes(){
	clientRoutes := s.engine.Group("/v1/client")


	routes.ClientRoutes(clientRoutes)

}

func (s *Server) Run() {
	log.Println("Server running on " + s.httpAddr)
	s.engine.Run(s.httpAddr)
}
