package dingtalk

import (
	"github.com/liunian1004/dingtalk/httplib"
	"log"
	"fmt"
	"time"
)

/*
官方参考文档：
https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.8JWX03&treeId=257&articleId=105735&docType=1

发送文字
{
    "msgtype": "text",
    "text": {
        "content": "我就是我, 是不一样的烟火"
    },
    "at": {
        "atMobiles": [
            "156xxxx8827",
            "189xxxx8325"
        ],
        "isAtAll": false
    }
}

link 类型
{
    "msgtype": "link",
    "link": {
        "text": "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
        "title": "时代的火车向前开",
        "picUrl": "",
        "messageUrl": "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
    }
}

 */

// 具体的消息数字题
type MsgText struct {
	MsgType string `json:"msgtype"`
	Text    Text `json:"text"`
	At      At `json:"at"`
}

// 文本内容主体
type Text struct {
	Content string `json:"content"`
}

// @人的操作
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool `json:"isAtAll"`
}

type Btns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// link类型
type MsgLink struct {
	MsgType string `json:"msgtype"`
	Link    Link `json:"link"`
}

// link 类型详细
type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

// 整体跳转ActionCard类型
type MsgActionCard struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

// 整体跳转ActionCard类型
type ActionCard struct {
	Text           string `json:"text"`
	Tile           string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

type MsgActionCardBtns struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCardBtns `json:"actionCard"`
}

// 独立跳转ActionCard类型
type ActionCardBtns struct {
	Text           string `json:"text"`
	Tile           string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []Btns `json:"btns"`
}

// FeedCard类型
type MsgFeedCard struct {
	MsgType      string `json:"msgtype"`
	FeedCardLink FeedCardLink `json:"feedCard"`
}

type FeedCardLink struct {
	Links []Link `json:"links"`
}

func (d *DingTalk) SendText(content string) error {
	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = content

	req, err := httplib.Post(d.GetUrl()).JSONBody(msgText)
	if err != nil {
		log.Println(err)
	}
	str, err := req.String()
	log.Println(str)
	return err
}

func (d *DingTalk) SendLog(content string) error {
	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = fmt.Sprintf("[%s]     [%s]     %s", time.Now().Format("2006-01-02 15:04:05"), GetIP(), content)

	req, err := httplib.Post(d.GetUrl()).JSONBody(msgText)
	if err != nil {
		log.Println(err)
	}
	bytes, err := req.Bytes()
	if err != nil {
		return err
	}
	e := NewDingError(bytes)
	if e.Errcode != 0 {
		// retry once
		bytes, err = req.Bytes()
		if err != nil {
			return err
		}
		e = NewDingError(bytes)
		if e.Errcode != 0 {
			return fmt.Errorf("发送失败！")
		}
	}
	return err
}
