package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var DB *gorm.DB

type APIResponse struct {
	Code int                    `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

func (r *APIResponse) ToBytes() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

func NewErrorResponse(code int, errMsg string) *APIResponse {
	return &APIResponse{
		Code: code,
		Msg:  errMsg,
	}
}

func NewErrorCodeResponse(code int) *APIResponse {
	return &APIResponse{
		Code: code,
		Msg:  ErrorCodes[code],
	}
}

func NewResponse(code int, data map[string]interface{}) *APIResponse {
	return &APIResponse{
		Code: code,
		Msg:  "ok",
		Data: data,
	}
}

func ConfigRouter() http.Handler {
	router := httprouter.New()

	// static
	router.Handler("GET", "/", http.FileServer(http.Dir("./webpages/web")))
	router.ServeFiles("/static/*filepath", http.Dir("./webpages/web/static"))

	// auth: login/logout
	router.POST("/login", loginHandler)
	router.POST("/logout", logoutHandler)

	//admin: handlers
	configAccountRoutes(router)
	configPostRoutes(router)
	configDCIRoutes(router)

	return router
}
