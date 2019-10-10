package routers

import (
	"fmt"
	"log"
	"net/http"
)

// IndexRoute = is '/'
// Передача параметров в template
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	user := ses.GetUserData(r)
	message := fmt.Sprintf("%s you are %s!", "Hello", user.Role)
	w.Write([]byte(message))

}

// Something - ...
func Something(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, voc_user, conroll_date FROM aggregate_skud order by 3 desc")
	if err != nil {
		log.Fatal(err)
	}
	data := []SKUD{}
	for rows.Next() {
		bk := new(SKUD)
		err := rows.Scan(&bk.ID, &bk.VocationUser, &bk.Date)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, *bk)
	}
	Jsonify(w, data, 200)
}
