package utils

import (
	"fmt"
	"crypto/md5"
)

// Jiami 加密
func Jiami(pwd *string,username *string)(string){
	h := md5.New()
	h.Write([]byte(*pwd))
	password := fmt.Sprintf("%x", h.Sum(nil))
	h.Write([]byte(password + *username))
	password = fmt.Sprintf("%x", h.Sum(nil))
	return password
}
