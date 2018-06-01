package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

func getInt(r *http.Request, name string, defaultValue int64) int64 {
	p := r.FormValue(name)
	if p == "" {
		return defaultValue
	}
	n, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return defaultValue
	}
	return n
}
