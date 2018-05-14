package api

import (
	"net/http"
	"strconv"
)

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
