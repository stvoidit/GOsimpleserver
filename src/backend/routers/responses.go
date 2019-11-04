package routers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
)

// Jsonify - отправка json response
func Jsonify(w http.ResponseWriter, i interface{}, code int) {
	switch i.(type) {
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		encoder := json.NewEncoder(w)
		encoder.Encode(i)
	case string:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(code)
		w.Write([]byte(i.(string)))
	case []byte:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(code)
		w.Write(i.([]byte))
	}
}

// JSONLoad - преобразование json в структуру
func JSONLoad(req *http.Request, i interface{}) (interface{}, error) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// TemplatesMap - ...
type TemplatesMap struct {
	sync.RWMutex
	template map[string][]byte
}

var tmp = TemplatesMap{template: make(map[string][]byte)}

// RegistrateTemplates - ...
func RegistrateTemplates(tmpPath string) {
	files, err := ioutil.ReadDir(tmpPath)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		fullpath := path.Join(tmpPath, file.Name())
		filename := strings.TrimSuffix(file.Name(), ".html")
		bfile, err := ioutil.ReadFile(fullpath)
		if err != nil {
			panic(err)
		}
		tmp.template[filename] = bfile
	}
}

// RenderTemplate - ...
func RenderTemplate(w http.ResponseWriter, name string) {
	tmp.RLock()
	defer tmp.RUnlock()
	w.Write(tmp.template[name])
}
