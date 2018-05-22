package handler

import (
	"bufio"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/weibocom/ipc/client"
	"github.com/weibocom/ipc/web2/model"
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

	if uid < 0 || uid >= 10000000000 {
		resp := NewErrorCodeResponse(40001000)
		w.Write(resp.ToBytes())
		return
	}
	if company == "" {
		resp := NewErrorCodeResponse(40001001)
		w.Write(resp.ToBytes())
		return
	}

	priKey, pubKey, err := service.RegisterAccount(company, uid)

	var resp *APIResponse
	if err != nil {
		if strings.Contains(err.Error(), "likely a uniqueness") {
			resp = NewErrorCodeResponse(40001002)
		} else {
			resp = NewErrorResponse(500, err.Error())
		}
	} else {

		user := &model.User{
			ID:         uid,
			Company:    company,
			PrivateKey: priKey,
			PublicKey:  pubKey,
		}

		if err != nil {
			resp = NewErrorResponse(500, err.Error())
			w.Write(resp.ToBytes())
			return
		}
		data := map[string]interface{}{"user": user}
		resp = NewResponse(200, data)

	}

	w.Write(resp.ToBytes())
}

// 分页查询所有的账号
func queryAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	company := r.FormValue("company")
	if company == "" {
		resp := NewErrorCodeResponse(40001001)
		w.Write(resp.ToBytes())
		return
	}

	page := getInt(r, "page", 1)
	pagesize := getInt(r, "pagesize", 20)

	uid := getInt(r, "uid", -1)

	users, err := service.GetUsers(company, int(page), int(pagesize), uid)

	var count int64

	if uid == -1 {
		count, err = service.UserCount()
	} else {
		count = int64(len(users))
	}

	var resp *APIResponse
	if err != nil {
		resp = NewErrorResponse(500, err.Error())
	} else {
		data := map[string]interface{}{"count": count, "users": users}
		resp = NewResponse(200, data)
	}

	w.Write(resp.ToBytes())
}

// 批量创建账户
func batchCreateAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		resp := NewErrorResponse(500, err.Error())
		w.Write(resp.ToBytes())
		return
	}

	company := r.FormValue("company")
	if company == "" {
		resp := NewErrorCodeResponse(40001001)
		w.Write(resp.ToBytes())
		return
	}

	accFile, header, err := r.FormFile("accounts_file")
	if err != nil {
		resp := NewErrorResponse(500, err.Error())
		w.Write(resp.ToBytes())
		return
	}
	defer accFile.Close()

	if filepath.Ext(header.Filename) != ".csv" {
		resp := NewErrorCodeResponse(40001010)
		w.Write(resp.ToBytes())
		return
	}

	var wrongFormatNum int
	var existAccountNum int
	var succeedNum int
	var errNum int

	scanner := bufio.NewScanner(accFile)
	for scanner.Scan() {
		// 因为只能上传本公司的用户，所以不需要第一列的公司名，

		fields := strings.Split(scanner.Text(), ",")
		uid, err := strconv.ParseInt(fields[len(fields)-1], 10, 64)
		if err != nil || uid < 0 || uid >= 10000000000 {
			wrongFormatNum++
			log.Printf("wrong format for uid: %d", uid)
			continue
		}

		_, _, err = service.RegisterAccount(company, uid)
		if err != nil {
			if err == client.ErrAccountAlreadyExist {
				existAccountNum++
			} else {
				errNum++
			}
		} else {
			succeedNum++
		}
	}

	data := map[string]interface{}{"wrong_format": wrongFormatNum,
		"existed":   existAccountNum,
		"succeeded": succeedNum,
		"failed":    errNum}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}

func accountDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	acc := ps.ByName("account")
	uid, err := strconv.ParseInt(acc, 10, 64)
	if err != nil || uid < 0 || uid >= 10000000000 {
		resp := NewErrorCodeResponse(40001000)
		w.Write(resp.ToBytes())
		return
	}

	company := r.FormValue("company")

	if uid < 0 || uid >= 10000000000 {
		resp := NewErrorCodeResponse(40001000)
		w.Write(resp.ToBytes())
		return
	}
	if company == "" {
		resp := NewErrorCodeResponse(40001001)
		w.Write(resp.ToBytes())
		return
	}

	user, err := service.GetUser(company, uid)
	if err != nil {
		resp := NewErrorCodeResponse(40001021)
		w.Write(resp.ToBytes())
		return
	}

	data := map[string]interface{}{"user": user}
	resp := NewResponse(200, data)
	w.Write(resp.ToBytes())
}
