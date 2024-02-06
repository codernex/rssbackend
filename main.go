package main

import (
	"database/sql"
	"github.com/codernex/rssbackend/internal/database"
	"github.com/codernex/rssbackend/utils"
	_ "github.com/codernex/rssbackend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occured finding .env %v", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalf("Error occured finding port")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalf("Error occured finding DB_URL")
	}
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Cant't connect to database")
	}
	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}
	router := chi.NewRouter()
	corsOpt := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	router.Use(corsOpt.Handler)
	router.Use(middleware.CleanPath)

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/healthz", handlerReadiness)
	v1Router.Get("/err", func(writer http.ResponseWriter, request *http.Request) {
		utils.RespondWithErr(writer, 500, "Some error")
	})
	//***** User Routes Start ******

	v1Router.Route("/users", func(r chi.Router) {
		r.Use(utils.IsAuthenticated)
		r.Post("/", apiCfg.handlerCreateUser)
		r.Get("/", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	})
	//***** User Routes End ******

	//***** Feed Routes Start ******
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds/{userId}", apiCfg.middlewareAuth(apiCfg.handlerGetFeedByUser))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	router.Mount("/v1", v1Router)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Error occured on server %v", err)
	}
}
