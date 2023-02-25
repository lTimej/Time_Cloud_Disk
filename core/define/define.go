package define

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var EmailPassword = os.Getenv("EmailPassword")

const CodeLen = 6

var CodeString = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const CodeExpire = 300

var TokenKey = []byte("cloud_disk")

type TokenClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

// MinIOAccessKeyID MinIO 配置
var MinIOAccessKeyID = "I0z4bcvvEKRJPA9F"
var MinIOAccessSecretKey = "sfMosGYCyN0xOERWwnvowfNLcTHASXyJ"
var MinIOEndpoint = "172.20.16.20:9000"
var MinIOBucket = "cloud-disk"

// 分页
var PageSize int = 10
