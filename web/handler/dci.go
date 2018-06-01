package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	con "github.com/weibocom/ipc/content"
	"github.com/weibocom/ipc/web/service"
)

// 版权鉴定 Digital Copyright Identification

func configDCIRoutes(router *httprouter.Router) {
	router.GET("/dci/content", auth(comparePost))
	router.GET("/dci/text", auth(compareText))
}

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
	uid1 := getInt(r, "src_uid", -1)
	mid1 := getInt(r, "src_mid", -1)
	company1 := r.FormValue("src_company")

	uid2 := getInt(r, "dst_uid", -1)
	mid2 := getInt(r, "dst_mid", -1)
	company2 := r.FormValue("dst_company")

	if !validateUIDMsgID(w, company1, uid1, mid1) || !validateUIDMsgID(w, company2, uid2, mid2) {
		return
	}

	post1, err := service.GetContentByMsgID(company1, uid1, mid1)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	post2, err := service.GetContentByMsgID(company2, uid2, mid2)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	s := con.Similarity(post1.Content, post2.Content) * 100

	data := map[string]interface{}{"similarity": fmt.Sprintf("%.2f", s),
		"src": post1.Content,
		"dst": post2.Content}

	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

// 根据DNA查询
func comparePostByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dna1 := r.FormValue("src_dna")
	if dna1 == "" {
		resp := NewErrorCodeResponse(40003005)
		w.Write(resp.ToBytes())
		return
	}

	dna2 := r.FormValue("dst_dna")
	if dna2 == "" {
		resp := NewErrorCodeResponse(40003005)
		w.Write(resp.ToBytes())
		return
	}

	post1, err := service.GetContentByDNA(dna1)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	post2, err := service.GetContentByDNA(dna2)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	s := con.Similarity(post1.Content, post2.Content) * 100

	data := map[string]interface{}{"similarity": fmt.Sprintf("%.2f", s),
		"src": post1.Content,
		"dst": post2.Content}

	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

func compareText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := r.FormValue("compareType")

	switch t {
	case "user":
		compareTextByUserPostID(w, r, ps)
	case "dna":
		compareTextByDNA(w, r, ps)
	default:
		resp := NewErrorCodeResponse(40003000)
		w.Write(resp.ToBytes())
	}
}

// 根据uid和mid查询
func compareTextByUserPostID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid1 := getInt(r, "src_uid", -1)
	mid1 := getInt(r, "src_mid", -1)
	company1 := r.FormValue("src_company")

	if !validateUIDMsgID(w, company1, uid1, mid1) {
		return
	}

	dstContent := r.FormValue("dst_content")
	if dstContent == "" {
		resp := NewErrorCodeResponse(40003006)
		w.Write(resp.ToBytes())
		return
	}

	post1, err := service.GetContentByMsgID(company1, uid1, mid1)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	s := con.Similarity(post1.Content, dstContent) * 100

	data := map[string]interface{}{"similarity": fmt.Sprintf("%.2f", s),
		"src": post1.Content,
		"dst": dstContent}

	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

// 根据DNA查询
func compareTextByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dna1 := r.FormValue("src_dna")
	if dna1 == "" {
		resp := NewErrorCodeResponse(40003005)
		w.Write(resp.ToBytes())
		return
	}

	dstContent := r.FormValue("dst_content")
	if dstContent == "" {
		resp := NewErrorCodeResponse(40003006)
		w.Write(resp.ToBytes())
		return
	}

	post1, err := service.GetContentByDNA(dna1)
	if err != nil {
		resp := NewErrorResponse(40003001, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	s := con.Similarity(post1.Content, dstContent) * 100

	data := map[string]interface{}{"similarity": fmt.Sprintf("%.2f", s),
		"src": post1.Content,
		"dst": dstContent}

	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())

}
