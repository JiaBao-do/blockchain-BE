package main

import (
	"log"
	"net/http"

	"github/project/blockchain/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {

	const (
		port = "8080"
	)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/test", handler.HandlerTest)
	v1Router.Get("/err", handler.HandlerErr)

	v1Router.Get("/start", handler.HandleStartBlockChain)
	v1Router.Get("/", handler.HandleGetBlockChain)
	v1Router.Post("/create", handler.HandleCreateBlock)

	router.Mount("/v1/blockchain", v1Router)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
