package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-package-integration/config"
	"github.com/mhdianrush/go-package-integration/controller"
	"github.com/mhdianrush/go-package-integration/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.Index).Methods(http.MethodGet)

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{
			"ok": true,
		})
		// or
		// result, _ := json.Marshal(map[string]bool{"ok": true})
		// w.Write(result)
	}).Methods(http.MethodGet)

	r.Use(middleware.LoggingMiddleware)

	logger := logrus.New()

	logger.Println("Server Running on Port 8080")

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
