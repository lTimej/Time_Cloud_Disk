package helper

import (
	"gopkg.in/gomail.v2"
	//"github.com/jordan-wright/email"
	// "crypto/tls"
	"liujun/Time_Cloud_Disk/core/define"
	// "fmt"
	"math/rand"
	// "net/smtp"
	"time"
	uuid "github.com/satori/go.uuid"
	"crypto/md5"
	"encoding/hex"
)

func SendEmail(to_Email string,code string)error{
	email_addr := "smtp.qq.com"
	email_user := "643163569@qq.com"
	email_port := 25
	// e := email.NewEmail()
	// e.From = fmt.Sprintf("Get %s",email_user)
	// e.To = []string{to_Email}
	// e.Subject = "验证码发送测试"
	// e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	// addr := fmt.Sprintf("%s:%d",email_addr,email_port)
	// err := e.SendWithTLS(addr, smtp.PlainAuth("", email_user, define.EmailPassword, email_addr),
	// 	&tls.Config{InsecureSkipVerify: true, ServerName: email_addr})
	// if err != nil {
	// 	return err
	// }
	// return nil
	m := gomail.NewMessage()
	m.SetHeader("From", email_user)
	m.SetHeader("To", to_Email)
	m.SetHeader("Subject", "用户注册验证码") 
	m.SetBody("text/html", "你的验证码为：<h1>" + code + "</h1>")
	d := gomail.NewDialer(email_addr, email_port, email_user, define.EmailPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func GetCode()string{
	b := make([]rune,define.CodeLen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b{
		b[i] = define.CodeString[r.Intn(len(define.CodeString))]
	}
	return string(b)
}

func UUID() string {
	return uuid.NewV4().String()
}

func MD5(pwd string)string{
	h := md5.New()
	h.Write([]byte(pwd))
	str := h.Sum(nil)
	return hex.EncodeToString(str)
}