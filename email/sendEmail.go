package email

import (
	"log"
	"net/smtp"
)

func SendEmail() {
	//这一块要开启qq邮箱的三方登录密匙，qq邮箱设置就可以拿到
	auth := smtp.PlainAuth("", "1033573740@qq.com", "xxxxxxxxxxx", "smtp.qq.com")
	to := []string{"xxxxxxxx@qq.com"}
	msg := []byte("To: 1033573839@qq.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.qq.com:25", auth, "xxxxxxx@qq.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
