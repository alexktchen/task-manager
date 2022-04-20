package main

import (
	"net/http"
	"time"

	"github.com/alexktchen/task-manager/routers"
)

func main() {
	routersInit := routers.Init()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
