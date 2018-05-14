package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var DB *gorm.DB

func ConfigRouter() http.Handler {
	router := httprouter.New()

	// static
	router.Handler("GET", "/", http.FileServer(http.Dir("./webpages/web")))
	router.ServeFiles("/static/*filepath", http.Dir("./webpages/web/static"))

	// legacy
	configBlogRouter(router)
	configAccountRouter(router)
	return router
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
