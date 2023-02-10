package logic

import (
	"context"

	"FirstProject/user-api/internal/svc"
	"FirstProject/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlerGetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlerGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlerGetUserInfoLogic {
	return &HandlerGetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlerGetUserInfoLogic) HandlerGetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	resp = &types.UserInfoResp{
		UserId: req.UserId,
	}
	return resp, nil
}
