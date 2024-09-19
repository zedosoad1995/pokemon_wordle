package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func ConfigCors(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    })

    return c.Handler(handler)
} 