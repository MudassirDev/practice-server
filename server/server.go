package server

import "net/http"

func CreateServer() *http.Server {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
