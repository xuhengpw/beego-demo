package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ResponsePage struct {
	Errcode		int			`json:"errcode"`
	Errmsg		string		`json:"errmsg"`
	PageData	*PageData	`json:"data"`
}
type PageData struct {
	Count	int64		`json:"count"`
	Data	interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  interface{} `json:"errmsg"`
}
