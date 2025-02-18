package controllers

import (
	"net/http"

	"github.com/MudassirDev/practice-server/internal/database"
	api_models "github.com/MudassirDev/practice-server/models/api"
)

func CreateUserController(cfg api_models.APIConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dbUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{msg: 'There was an error while creating user'}"))
			return
		}
	}
}
