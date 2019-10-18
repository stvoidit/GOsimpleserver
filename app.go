package main

import (
	"net/http"

	"./application"
)

func main() {
	app := application.App
	// logfile, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// defer logfile.Close()
	// handlers.LoggingHandler(logfile, r)
	http.ListenAndServe("0.0.0.0:9000", app.Router)
}
