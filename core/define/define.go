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
var MinIOAccessKeyID = "WHPITDVZRR317H7TTGBN"
var MinIOAccessSecretKey = "mDphi8CpUVHkZ4noRY9fTx5GawtXaRs9o1elfWfa"
var MinIOEndpoint = "192.168.70.129:9000"
var MinIOBucket = "cloud-disk"