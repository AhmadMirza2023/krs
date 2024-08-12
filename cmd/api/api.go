package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/AhmadMirza2023/krs/course"
	"github.com/AhmadMirza2023/krs/service/user"
	"github.com/AhmadMirza2023/krs/spc"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Course
	courseStore := course.NewStore(s.db)
	courseHandler := course.NewHandler(courseStore)
	courseHandler.RegisterRoutes(subrouter)

	// SPC
	spcStore := spc.NewStore(s.db)
	spcHandler := spc.NewHandler(spcStore)
	spcHandler.RegisterRoutes(subrouter)

	log.Println("Listening on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
