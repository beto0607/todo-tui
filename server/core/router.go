package core

import "net/http"

func InitRouting() *http.ServeMux {
	router := http.NewServeMux()

	return router
}
