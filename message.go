package gox

import (
	"fmt"
)

type (
	// Code 消息代码
	Code int

	// CodeMessage 带消息代码的消息
	CodeMessage struct {
		// Code 消息代码
		Code Code `json:"code"`
		// Title 消息标题
		// 通常是模板字符串
		Title *string `json:"title"`
		// Text 消息内容
		// 通常是模板字符串
		Text *string `json:"text,omitempty"`
	}
)

// NewCodeMessage 创建消息
func NewCodeMessage(code Code, title, text string) *CodeMessage {
	return &CodeMessage{
		Code:  code,
		Title: &title,
		Text:  &text,
	}
}

func (cm CodeMessage) ToCode() Code {
	return cm.Code
}

func (cm CodeMessage) String() string {
	return fmt.Sprintf("code=%d, title=%s, text=%s", cm.Code, *cm.Title, *cm.Text)
}

func (cm CodeMessage) ParseTitle(params ...interface{}) CodeMessage {
	title := fmt.Sprintf(*cm.Title, params...)
	cm.Title = &title

	return cm
}

func (cm CodeMessage) ParseText(params ...interface{}) CodeMessage {
	text := fmt.Sprintf(*cm.Text, params...)
	cm.Text = &text

	return cm
}
