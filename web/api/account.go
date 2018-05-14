package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/weibocom/ipc/web/service"
)

type APIResponse struct {
	RetCode int                    `json:"retCode,omitempty"`
	Msg     string                 `json:"msg,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (r *APIResponse) ToBytes() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

func configAccountRouter(router *httprouter.Router) {
	router.GET("/account/count.json", accountCount)
	router.GET("/register.json", registerAccount)
}

func accountCount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	count, err := service.GetUserCount()
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
			Data:    map[string]interface{}{"count": count},
		}
	}
	w.Write(resp.ToBytes())
}

func registerAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	uid, _ := strconv.ParseInt(r.FormValue("uid"), 10, 64)
	company := r.FormValue("company")

	err := service.RegisterIPAccount(company, uid)

	var resp *APIResponse
	if err != nil {
		if strings.Contains(err.Error(), "likely a uniqueness") {
			resp = &APIResponse{
				RetCode: 400,
				Msg:     err.Error(),
			}
		} else {
			resp = &APIResponse{
				RetCode: 500,
				Msg:     err.Error(),
			}
		}
	} else {
		resp = &APIResponse{
			RetCode: 200,
			Msg:     "ok",
			Data:    map[string]interface{}{"uid": uid, "company": company},
		}
	}

	w.Write(resp.ToBytes())
}
