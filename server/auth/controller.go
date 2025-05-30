package auth

import "net/http"

func InitRouting() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /auth/privacy", AuthPrivacy)
	router.HandleFunc("GET /auth/terms", AuthTerms)

	return router

}
