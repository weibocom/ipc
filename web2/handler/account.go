package handler

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/weibocom/ipc/web2/service"
)

// 为web提供的账号相关的rest api

func configAccountRoutes(router *httprouter.Router) {
	router.GET("/counts/account", auth(accountCount))
	router.POST("/accounts", auth(createAccount))
	router.GET("/accounts", auth(queryAccount))
	router.GET("/accounts/:account", auth(accountDetail))
}

func accountCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	count, err := service.UserCount()
	var resp *APIResponse
	if err != nil {
		resp = NewErrorResponse(500, err.Error())
	} else {
		resp = NewResponse(200, map[string]interface{}{"count": count})
	}
	w.Write(resp.ToBytes())
}

func createAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.FormValue("batch") == "true" || r.FormValue("batch") == "1" {
		batchCreateAccounts(w, r, ps)
		return
	}

	uid := getInt(r, "uid", -1)
	company := r.FormValue("company")

	if uid < 0 && uid >= 10000000000 {
		resp := NewErrorCodeResponse(40001000)
		w.Write(resp.ToBytes())
		return
	}
	if company == "" {
		resp := NewErrorCodeResponse(40001001)
		w.Write(resp.ToBytes())
		return
	}

	err := service.RegisterAccount(company, uid)

	var resp *APIResponse
	if err != nil {
		if strings.Contains(err.Error(), "likely a uniqueness") {
			resp = NewErrorCodeResponse(40001002)
		} else {
			resp = NewErrorResponse(500, err.Error())
		}
	} else {
		// TODO: private key and public key
		data := map[string]interface{}{"uid": uid, "company": company}
		resp = NewResponse(200, data)

	}

	w.Write(resp.ToBytes())
}

// TODO: 分页查询所有的账号
func queryAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// TODO: 批量创建账户
func batchCreateAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// TODO: 显示账户详情
func accountDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
