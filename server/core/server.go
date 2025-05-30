package core

import "net/http"

func InitServer(host string, port string) *http.Server {

	router := InitRouting()

	serverAddress := host + ":" + port

	server := http.Server{
		Addr:         serverAddress,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		Handler:      router,
	}
	return &server
}
