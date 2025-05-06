package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/houseofaragon/go_project/internal/app"
	"github.com/houseofaragon/go_project/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	// instantiate
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// http.HandleFunc(
	// 	"/health", app.HealthCheck,
	// )
	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf(
		"We are running app on port %d\n", port,
	)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
