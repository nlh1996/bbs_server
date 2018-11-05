package utils

import (
	"fmt"
	"crypto/md5"
)

// Jiami 加密
func Jiami(arg1 *string,arg2 *string)(string){
	h := md5.New()
	h.Write([]byte(*arg1))
	result := fmt.Sprintf("%x", h.Sum(nil))
	h.Write([]byte(result + *arg2))
	result = fmt.Sprintf("%x", h.Sum(nil))
	return result
}
