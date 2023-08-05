package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/diegom0ta/goauth/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.RenderPage)
	router.HandleFunc("/google-sso", controllers.GoogleSignOn)
	router.HandleFunc("/callback", controllers.Callback)

	srv := &http.Server{
		Handler: router,

		// you can replace 8007 with any other available port
		Addr: "127.0.0.1:8007",

		// enforces timeouts for created servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("listening on port: %s\n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
