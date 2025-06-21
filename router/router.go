package router

import (
	"net/http"

	"github.com/gabrielmartins0312/crud-go-with-api-and-db/handler"
)

func LoadRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAllUsers(w, r)
		case http.MethodPost:
			handler.CreateUser(w, r)
		case http.MethodPut:
			handler.UpdateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
