package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	ipcmodel "github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/web/model"
	"github.com/weibocom/ipc/web/service"

	con "github.com/weibocom/ipc/content"
)

func configBlogRouter(router *httprouter.Router) {
	router.GET("/compare.json", compareBlog)
	router.GET("/submit.json", addBlog)
	router.GET("/search.json", searchBlog)

	router.GET("/ts/fetch.json", fetchStatus)
}

func compareBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	srcAuthor := getInt(r, "srcUid", 0)
	srcMid := getInt(r, "srcMid", 0)
	srcCompany := r.FormValue("srcCompany")
	if srcCompany == "" {
		srcCompany = "weibo"
	}
	srcContent := r.FormValue("srcContent")
	if srcContent == "" {
		post, err := service.GetContentByMsgID(srcCompany, srcAuthor, srcMid)
		if err != nil {
			resp := &APIResponse{
				RetCode: 400,
				Msg:     err.Error(),
			}
			w.Write(resp.ToBytes())
			return
		}

		srcContent = post.Content
	}

	dstAuthor := getInt(r, "dstUid", 0)
	dstMid := getInt(r, "dstMid", 0)
	dstCompany := r.FormValue("dstCompany")
	if dstCompany == "" {
		dstCompany = "weibo"
	}
	post, err := service.GetContentByMsgID(dstCompany, dstAuthor, dstMid)
	if err != nil {
		resp := &APIResponse{
			RetCode: 400,
			Msg:     err.Error(),
		}
		w.Write(resp.ToBytes())
		return
	}

	dstContent := post.Content

	s := con.Similarity(srcContent, dstContent) * 100

	resp := &APIResponse{
		RetCode: 200,
		Msg:     "ok",
		Data: map[string]interface{}{"similarity": fmt.Sprintf("%.2f", s),
			"srcContent": srcContent,
			"dstContent": dstContent},
	}

	w.Write(resp.ToBytes())
}

func addBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	action := r.FormValue("action")
	if action != "write" {
		searchBlog(w, r, ps)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	uid := getInt(r, "uid", 0)
	mid := getInt(r, "mid", 0)
	company := r.FormValue("company")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var resp *APIResponse

	dna, err := service.SendBlog(uid, mid, company, title, content, time.Now().UnixNano()/1e6)
	if err != nil {
		resp := &APIResponse{
			RetCode: 500,
			Msg:     err.Error(),
		}
		w.Write(resp.ToBytes())
		return
	}

	post, err := service.GetContent(company, uid, dna.ID())
	if err != nil {
		resp := &APIResponse{
			RetCode: 400,
			Msg:     err.Error(),
		}
		w.Write(resp.ToBytes())
		return
	}

	viewDiscussion := &model.ViewDiscussion{}
	if post != nil {
		viewDiscussion.AuthorName = post.Author
		viewDiscussion.BlockTitle = post.Title
		viewDiscussion.Content = post.Content
		viewDiscussion.BlockCreateTime = time.Now()
		viewDiscussion.Url = post.URI

		resp = &APIResponse{
			RetCode: 200,
			Msg:     "ok",
			Data:    map[string]interface{}{"blockInfo": viewDiscussion},
		}
	}

	w.Write(resp.ToBytes())
}

func searchBlog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	uid := getInt(r, "uid", 0)
	mid := getInt(r, "mid", 0)
	company := r.FormValue("company")

	post, err := service.GetContentByMsgID(company, uid, mid)
	if err != nil {
		resp := &APIResponse{
			RetCode: 400,
			Msg:     err.Error(),
		}
		w.Write(resp.ToBytes())
		return
	}

	var resp *APIResponse
	viewDiscussion := &model.ViewDiscussion{}
	if post != nil {
		viewDiscussion.AuthorName = post.Author
		viewDiscussion.BlockTitle = post.Title
		viewDiscussion.Content = post.Content
		viewDiscussion.BlockCreateTime = post.CreatedAt
		viewDiscussion.Url = post.URI

		resp = &APIResponse{
			RetCode: 200,
			Msg:     "ok",
			Data:    map[string]interface{}{"blockInfo": viewDiscussion},
		}
	}

	w.Write(resp.ToBytes())
}

func fetchStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	post, err := service.GetLatestMid()
	if post == nil {
		post = &ipcmodel.Post{}
	}
	var resp *APIResponse
	if err != nil {
		resp = &APIResponse{
			RetCode: 500,
			Msg:     err.Error(),
		}
	} else {
		resp = &APIResponse{
			RetCode: 200,
			Msg:     "ok",
			Data:    map[string]interface{}{"timestamp": post.CreatedAt},
		}
	}

	w.Write(resp.ToBytes())
}
