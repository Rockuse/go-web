package routes

import (
	"go-web/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()
	// //Route
	router.HandleFunc("/", handler.HelloHandler)
	router.HandleFunc("/fahmi", handler.HelloFahmi)
	router.HandleFunc("/ano", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Fahmi"))
	})
	router.HandleFunc("/anonymouse", handler.Anonymouse)
	router.HandleFunc("/product", handler.Product)
	return router
}
