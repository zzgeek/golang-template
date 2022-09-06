package model

type ResponseView struct {
	RespCode int         `json:"code"`
	RespMsg  string      `json:"message"`
	RespData interface{} `json:"respData"`
}
