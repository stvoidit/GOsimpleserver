package routers

import (
	"fmt"
	"net/http"
	// "../models"
)

// IndexRoute = is '/'
// Передача параметров в template
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("%s you are %s!", "Hello", "ADMIN")
	w.Write([]byte(message))

}
