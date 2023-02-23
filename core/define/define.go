package define

import (
	"os"
)

var EmailPassword = os.Getenv("EmailPassword")

const CodeLen = 6
var CodeString = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const CodeExpire = 300

