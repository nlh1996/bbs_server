package model

import "time"

// Thread 贴子结构
type Thread struct {
	TopStorey 		`json:"topStorey"`
	Reply1 				`json:"reply1"`
	Reply2 				`json:"reply2"`
	TID		string 	`json:"tid"`
}

// TopStorey .
type TopStorey struct {
	UID        string `json:"uid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	ReadNum    int32 	`json:"readNum"`
	Support    int32 	`json:"support"`
	ReplyNum   int32	`json:"replyNum"`
	CurrentTime time.Time`json:"createTime"`
	ImgList 	 []string	`json:"imgList"`
}	


// Reply1 .
type Reply1 []struct {
	UID				 string	`json:"uid"`
	Index			 int32	`json:"index"`
	Content    string	`json:"content"`
	CurrentTime time.Time`json:"createTime"`
} 

// Reply2 .
type Reply2 []struct {
	UID				 string	`json:"uid"`
	Index			 int32	`json:"index"`	
	Content    string	`json:"content"`
	CurrentTime time.Time`json:"createTime"`
}	