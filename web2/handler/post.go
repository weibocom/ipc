package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 为web提供的内存相关的rest api

func configPostRoutes(router *httprouter.Router) {
	router.GET("/posts", auth(queryPost))
	router.POST("/posts", auth(addPost))
}

// TODO: 根据uid和mid查询或者根据dna查询
func queryPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := r.FormValue("queryType")

	switch t {
	case "user":
		queryPostByUserPostID(w, r, ps)
	case "dna":
		queryPostByDNA(w, r, ps)
	default:
		resp := NewErrorCodeResponse(40002000)
		w.Write(resp.ToBytes())
	}
}

// 根据uid和mid查询
func queryPostByUserPostID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// 根据DNA查询
func queryPostByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func addPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
