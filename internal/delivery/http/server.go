package http

import (
	"net/http"

	repo "github.com/Bojong-Ireng/backend-server/internal/repository"
	"github.com/gorilla/mux"
)

func NewServer() http.Handler {
	router := mux.NewRouter()
	db := repo.GetConnection()

	addRoutes(
		router,
		db,
	)

	var handler http.Handler = router

	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)

	return handler
}
