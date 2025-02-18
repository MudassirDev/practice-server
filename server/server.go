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
	_ "github.com/lib/pq"
)

func CreateServer() *http.Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't read dotenv files")
	}
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
