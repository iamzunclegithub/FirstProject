package logic

import (
	"context"

	"FirstProject/user/internal/svc"
	"FirstProject/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{
		Message: "Hello " + req.Name + "!" + " This is a test message from go-zero.",
	}
	l.Logger.Infof("[UserLogic] Hello %s!", req.Name)
	return resp, nil
}
