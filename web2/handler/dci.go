package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// 版权鉴定 Digital Copyright Identification

func configDCIRoutes(router *httprouter.Router) {
	router.GET("/dci/resource", auth(comparePost))
	router.GET("/dci/text", auth(compareText))
	router.GET("/dci/cert", auth(cert))
}

// TODO:
func comparePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := r.FormValue("compareType")

	switch t {
	case "user":
		comparePostByUserPostID(w, r, ps)
	case "dna":
		comparePostByDNA(w, r, ps)
	default:
		resp := NewErrorCodeResponse(40003000)
		w.Write(resp.ToBytes())
	}
}

// 根据uid和mid查询
func comparePostByUserPostID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// 根据DNA查询
func comparePostByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// TODO:
func compareText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := r.FormValue("compareType")

	switch t {
	case "user":
		comparePostByUserPostID(w, r, ps)
	case "dna":
		comparePostByDNA(w, r, ps)
	default:
		resp := NewErrorCodeResponse(40003000)
		w.Write(resp.ToBytes())
	}
}

// 根据uid和mid查询
func compareTextByUserPostID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// 根据DNA查询
func compareTextByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// TODO: 生成鉴证证书
func cert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
