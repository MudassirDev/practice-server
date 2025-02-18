package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MudassirDev/practice-server/internal/database"
	Json "github.com/MudassirDev/practice-server/json"
	api_models "github.com/MudassirDev/practice-server/models/api"
)

func CreateUserController(cfg api_models.APIConfig) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestBodyParams struct {
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var Data RequestBodyParams
		if err := decoder.Decode(&Data); err != nil {
			Json.ResponseWithError(w, "error occured during reading request", http.StatusInternalServerError)
			return
		}

		dbUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Email:     Data.Email,
			FirstName: Data.FirstName,
			LastName:  Data.LastName,
		})

		if err != nil {
			Json.ResponseWithError(w, "error occured during creating user", http.StatusInternalServerError)
			return
		}

		Json.ResponseWithJson(w, dbUser, http.StatusOK)
	}
}
