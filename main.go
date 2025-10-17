package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"fireUp/internal/app"
	"fireUp/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go back-end server port")
	flag.Parse()

	app, err := app.NewApplciton()

	if err != nil {
		panic(err)
	}
	// Holy moly
	defer app.DB.Close()

	r := routes.SetUpRoutes(app)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Server is listening on port %d\n", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}

