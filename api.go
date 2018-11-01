package dingtalk

import (
	"fmt"
)

const Endpoint string = "https://oapi.dingtalk.com/robot/send?access_token=%s"

type DingTalk struct {
    *Context
}

func NewDingTalk(c *Context) *DingTalk {
	return &DingTalk{ Context: c }
}

func (d *DingTalk) GetUrl() string {
	return fmt.Sprintf(Endpoint, d.AccessToken)
}