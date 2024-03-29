package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rss-feed/internal/database"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB      *database.Queries
	Handler http.Handler
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiConfig := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiRouterV1 := chi.NewRouter()
	apiRouterV1.Get("/health", handlerReadiness)
	apiRouterV1.Get("/err", handlerError)

	apiRouterV1.Post("/users", apiConfig.handlerCreateUsers)
	apiRouterV1.Get("/users", apiConfig.middleWareAuth(apiConfig.handlerGetUser))

	apiRouterV1.Post("/feeds", apiConfig.middleWareAuth(apiConfig.handlerCreateFeed))
	apiRouterV1.Get("/feeds", apiConfig.handlerGetFeeds)

	apiRouterV1.Get("/feed_follows", apiConfig.middleWareAuth(apiConfig.handlerGetFeedFollows))
	apiRouterV1.Post("/feed_follows", apiConfig.middleWareAuth(apiConfig.handlerCreateFeedFollow))
	apiRouterV1.Delete("/feed_follows/{feedFollowID}", apiConfig.middleWareAuth(apiConfig.handlerDeleteFeedFollow))

	apiRouterV1.Get("/posts", apiConfig.middleWareAuth(apiConfig.handlerGetPosts))

	router.Mount("/v1", apiRouterV1)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	const collectioncConcurrency = 10
	const collectionInterval = time.Minute

	go startScraping(dbQueries, collectioncConcurrency, collectionInterval)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
