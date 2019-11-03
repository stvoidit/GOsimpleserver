package main

import (
	"net/http"
	"os"

	"./application"
	"github.com/gorilla/handlers"
)

func main() {
	app := application.Start()
	logfile, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	http.ListenAndServe("127.0.0.1:9000", handlers.LoggingHandler(logfile, app))
	// http.ListenAndServe("127.0.0.1:9000", handlers.LoggingHandler(os.Stdout, app))
}
