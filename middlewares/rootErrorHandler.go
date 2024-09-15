package middlewares

import (
	"fmt"
	"net/http"

	route_types "github.com/zedosoad1995/pokemon-wordle/routes/types"
)

func ErrorHandler(next route_types.RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		if err := next(w, r); err != nil {
			fmt.Printf("Something went wrong: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
