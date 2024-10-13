package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Router *chi.Mux
	// Db, config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers(dbxhandler *tasksDefinitions) {
	// Mount all Middleware here
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Timeout(time.Millisecond * 40))
	// Mount all handlers here

	s.Router.Get("/getall", dbxhandler.ApiGetAllTasks)
	s.Router.Get("/getstats", dbxhandler.ApiGetStats)
	s.Router.Get("/findone", dbxhandler.ApiFindOneTasks)

}

