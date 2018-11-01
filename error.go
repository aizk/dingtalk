package dingtalk

import "encoding/json"

type DingError struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
}

func NewDingError(bytes []byte) *DingError {
	e := &DingError{}
	json.Unmarshal(bytes, e)
	return e
}