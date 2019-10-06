package routers

import (
	"fmt"
	"net/http"
	// "../models"
)

// IndexRoute = is '/'
// Передача параметров в template
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	user := ses.GetUserData(r)
	message := fmt.Sprintf("%s you are %s!", "Hello", user.Role)
	w.Write([]byte(message))

}
