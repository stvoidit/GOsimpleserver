package main

import (
	"net/http"
	"os"

	"./application"
	"github.com/gorilla/handlers"
)

func main() {
	app := application.App
	logfile, _ := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer logfile.Close()
	http.ListenAndServe("0.0.0.0:9000", handlers.LoggingHandler(logfile, app.Router))
}
