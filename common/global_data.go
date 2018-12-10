package common

import "bbs_server/model"

// TokenMap .
var (
	TokenMap  map[string]string //保存token
	PostsPool *[]model.Post    //贴子缓存池
	BlackList *[]model.BlackName //黑名单
)

func init() {
	TokenMap = make(map[string]string)
	PostsPool = &[]model.Post{}
	BlackList = &[]model.BlackName{}
}
