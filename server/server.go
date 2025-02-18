package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/MudassirDev/practice-server/controllers"
	"github.com/MudassirDev/practice-server/internal/database"
	"github.com/MudassirDev/practice-server/models/api"
	"github.com/joho/godotenv"
)

func CreateServer() *http.Server {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(conn)
	apiCfg := api_models.APIConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/register", controllers.CreateUserController(apiCfg))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
