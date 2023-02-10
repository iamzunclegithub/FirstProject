// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `json:"userId" in:"path" name:"userId" desc:"用户id"`
}

type UserInfoResp struct {
	UserId   int64  `json:"userId" in:"body" name:"userId" desc:"用户id"`
	NickName string `json:"nickName" in:"body" name:"nickName" desc:"昵称"`
}
