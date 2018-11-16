package common

import "bbs_server/model"

// TokenMap .
var (
	TokenMap  map[string]string //保存token
	PostsPool *[]model.Post    //贴子缓存池
)

func init() {
	TokenMap = make(map[string]string)
	PostsPool = &[]model.Post{}
}
