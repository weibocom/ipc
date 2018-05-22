package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/weibocom/ipc/web2/service"
)

// 为web提供的内存相关的rest api

func configPostRoutes(router *httprouter.Router) {
	router.GET("/counts/post", auth(postCount))
	router.GET("/last_post", auth(lastPost))
	router.GET("/posts", auth(queryPost))
	router.GET("/account_posts", auth(queryAccountPost))
	router.POST("/posts", auth(addPost))
}

func postCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	count, err := service.PostCount()
	var resp *APIResponse
	if err != nil {
		resp = NewErrorResponse(500, err.Error())
	} else {
		resp = NewResponse(200, map[string]interface{}{"count": count})
	}
	w.Write(resp.ToBytes())
}

func lastPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post, err := service.GetLatestPost()
	var resp *APIResponse
	if err != nil {
		resp = NewErrorResponse(500, err.Error())
	} else {
		resp = NewResponse(200, map[string]interface{}{"timestamp": post.CreatedAt})
	}
	w.Write(resp.ToBytes())
}

// 根据uid和mid查询或者根据dna查询
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
	uid := getInt(r, "uid", -1)
	mid := getInt(r, "mid", -1)
	company := r.FormValue("company")

	if !validateUIDMsgID(w, company, uid, mid) {
		return
	}

	post, err := service.GetContentByMsgID(company, uid, mid)
	if err != nil {
		resp := NewErrorCodeResponse(40002004)
		w.Write(resp.ToBytes())
		return
	}

	data := map[string]interface{}{"post": post}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

func validateUIDMsgID(w http.ResponseWriter, company string, uid, mid int64) bool {
	if company == "" {
		resp := NewErrorCodeResponse(40002001)
		w.Write(resp.ToBytes())
		return false
	}

	if uid == -1 {
		resp := NewErrorCodeResponse(40002002)
		w.Write(resp.ToBytes())
		return false
	}

	if mid == -1 {
		resp := NewErrorCodeResponse(40002003)
		w.Write(resp.ToBytes())
		return false
	}

	return true
}

// 根据DNA查询
func queryPostByDNA(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dna := r.FormValue("dna")

	if dna == "" {
		resp := NewErrorCodeResponse(40002005)
		w.Write(resp.ToBytes())
		return
	}

	post, err := service.GetContentByDNA(dna)
	if err != nil {
		resp := NewErrorCodeResponse(40002004)
		w.Write(resp.ToBytes())
		return
	}

	data := map[string]interface{}{"post": post}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

func queryAccountPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := getInt(r, "uid", -1)
	company := r.FormValue("company")

	if company == "" {
		resp := NewErrorCodeResponse(40002001)
		w.Write(resp.ToBytes())
		return
	}

	if uid == -1 {
		resp := NewErrorCodeResponse(40002002)
		w.Write(resp.ToBytes())
		return
	}

	page := getInt(r, "page", 1)
	pagesize := getInt(r, "pagesize", 20)

	posts, postCount, err := service.GetUserPosts(company, uid, int(page), int(pagesize))
	if err != nil {
		resp := NewErrorResponse(500, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	data := map[string]interface{}{"post_count": postCount, "posts": posts}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

func addPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := getInt(r, "uid", -1)
	mid := getInt(r, "mid", -1)
	company := r.FormValue("company")
	contentType := r.FormValue("contentType")
	content := r.FormValue("content")
	title := r.FormValue("title")

	if !validateUIDMsgID(w, company, uid, mid) {
		return
	}

	if content == "" {
		resp := NewErrorCodeResponse(40002006)
		w.Write(resp.ToBytes())
		return
	}

	if title == "" {
		title = fmt.Sprintf("%s-%d-%d", company, uid, mid)
	}
	_ = contentType

	dna, err := service.AddPost(company, uid, mid, title, content, time.Now().UnixNano()/1e6)
	if err != nil {
		resp := NewErrorResponse(500, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	post, err := service.GetContentByDNA(dna.ID())
	if err != nil {
		resp := NewErrorResponse(500, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	data := map[string]interface{}{"post": post}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}
