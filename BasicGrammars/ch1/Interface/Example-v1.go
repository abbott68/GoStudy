package main

import "fmt"

// 定义接口
type Messenger interface {
	SendMessnsger()
}

// 定义结构体
type Email struct {
	Recipient string
	Messeng   string
}

type SMS struct {
	Recipient string
	Messeng   string
}

// 定义实现接口方法
func (e Email) SendMessnsger() { //发送邮件
	fmt.Printf("Sending email to %s: %s\n", e.Recipient, e.Messeng)
}
func (s SMS) SendMessnsger() { //发送信息
	fmt.Printf("Sending SMS to %s: %s\n", s.Recipient, s.Messeng)
}
func main() {
	email := Email{
		Recipient: "abbott@126.com",
		Messeng:   "hello, this is an email",
	}
	sms := SMS{
		Recipient: "18301002516",
		Messeng:   "Hello, hi",
	}
	// 使用接口类型的变量调用方法
	var m Messenger
	m = email
	m.SendMessnsger()

	m = sms
	m.SendMessnsger()
}
