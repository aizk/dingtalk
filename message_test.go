package dingtalk

import (
	"testing"
	"time"
)

func TestDingTalk_SendText(t *testing.T) {
	d := NewDingTalk(&Context{
		AccessToken: "4f5fe91bc652973b278dbde1ed84821d9af159f40a39619d21423e33e83e7be7",
	})

	for {
		d.SendLog("exec script x success.")
		time.Sleep(5 * time.Second)
	}
}
