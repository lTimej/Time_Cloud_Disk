package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"
	"liujun/Time_Cloud_Disk/core/define"
	"math/rand"
	"crypto/md5"
	"context"
	"net/http"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"time"
	"path"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SendEmail(to_Email string, code string) error {
	email_addr := "smtp.qq.com"
	email_user := "643163569@qq.com"
	email_port := 25
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


func MinIOUpload(r *http.Request)(string,error) {
	ctx := context.Background()
	// Initialize minio client object.
	minioClient, err := minio.New(define.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(define.MinIOAccessKeyID, define.MinIOAccessSecretKey, ""),
	})
	if err != nil {
		return "",err
	}
	// Make a new bucket called mymusic.
	// bucketName := "mymusic"
	// location := "us-east-1"

	// err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	// if err != nil {
	// 	// Check to see if we already own this bucket (which happens if you run this twice)
	// 	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	// 	if errBucketExists == nil && exists {
	// 		log.Printf("We already own %s\n", bucketName)
	// 	} else {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	// 	log.Printf("Successfully created %s\n", bucketName)
	// }

	// Upload the zip file
	file,fileHandler,err := r.FormFile("file")
	objectName := UUID() + path.Ext(fileHandler.Filename)
	contentType := "binary/octet-stream"
	// Upload the zip file with FPutObject
	_, err = minioClient.PutObject(ctx, define.MinIOBucket, objectName, file, fileHandler.Size,minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "",err
	}

	return define.MinIOBucket + "/" + objectName,nil
}