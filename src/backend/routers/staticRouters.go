package routers

import "net/http"

// MyVieos - ...
func MyVieos(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "MyVieos")
}
