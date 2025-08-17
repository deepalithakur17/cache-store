package routes

import (
	"gorm-gogo/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/store", controllers.SetCache).Methods("POST")
	r.HandleFunc("/getChache", controllers.GetCache).Methods("GET")
	return r
}
