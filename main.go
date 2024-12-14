package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TBVanderley4545/the-one-with-a-go-server/controller"
	_ "github.com/TBVanderley4545/the-one-with-a-go-server/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = 8080

// @title Test Go API
// @version 1.0
func main() {
	log.Printf("Server has started on port: %d\n", port)

	router := http.NewServeMux()

	router.HandleFunc("GET /healthz", controller.HealthCheck)
	router.HandleFunc("GET /swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
