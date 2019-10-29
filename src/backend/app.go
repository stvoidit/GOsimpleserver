package main

import (
	"net/http"
	"os"

	"./application"
	"github.com/gorilla/handlers"
)

func main() {
	app := application.App
	logfile, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	http.ListenAndServe("0.0.0.0:9000", handlers.LoggingHandler(logfile, app.Router))
	// http.ListenAndServe("0.0.0.0:9000", handlers.LoggingHandler(os.Stdout, app.Router))
}
