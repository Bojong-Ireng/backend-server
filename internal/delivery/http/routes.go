package http

import (
	"context"
	"log"
	"net/http"

	"github.com/Bojong-Ireng/backend-server/internal/domain"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func addRoutes(
	router *mux.Router,
	db *gorm.DB,
) {
	router.Handle("/test/postdb", handlePostDb(db))
}

func handlePostDb(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			encode(w, int(http.StatusMethodNotAllowed), "method not allowed")
			return
		}

		busData, err := decode[domain.Bus](r)
		if err != nil {
			encode(w, int(http.StatusBadRequest), "failed to decode fields. perhaps check your request body.")
		}

		ctx := context.Background()
		if err := gorm.G[domain.Bus](db).Create(ctx, &busData); err != nil {
			log.Printf("failed to store values into database: %v", err)
			encode(w, int(http.StatusInternalServerError), "server error, please report to our team so we can fix the issue")
			return
		}

		encode(w, http.StatusOK, "successfully posted to db")
	})
}
