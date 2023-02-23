package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
	//"github.com/jordan-wright/email"
	// "crypto/tls"
	"liujun/Time_Cloud_Disk/core/define"
	// "fmt"
	"math/rand"
	// "net/smtp"
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"time"
)

func SendEmail(to_Email string, code string) error {
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
	m.SetBody("text/html", "你的验证码为：<h1>"+code+"</h1>")
	d := gomail.NewDialer(email_addr, email_port, email_user, define.EmailPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func GetCode() string {
	b := make([]rune, define.CodeLen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = define.CodeString[r.Intn(len(define.CodeString))]
	}
	return string(b)
}

func UUID() string {
	return uuid.NewV4().String()
}

func MD5(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	str := h.Sum(nil)
	return hex.EncodeToString(str)
}

func GenToken(id int, identity, name string, second time.Duration) (string, error) {
	claim := define.TokenClaim{
		id,
		identity,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(second)),
			Issuer:    "cloud_disk", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(define.TokenKey)
}

func VerifyToken(token_str string) (*define.TokenClaim, error) {
	tc := new(define.TokenClaim)
	token, err := jwt.ParseWithClaims(token_str, tc, func(token *jwt.Token) (interface{}, error) {
		return define.TokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*define.TokenClaim); ok && token.Valid {
		return tc, nil
	}
	return nil, errors.New("非法token")
}
