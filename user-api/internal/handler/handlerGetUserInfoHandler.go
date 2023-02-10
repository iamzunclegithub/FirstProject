package handler

import (
	"net/http"

	"FirstProject/user-api/internal/logic"
	"FirstProject/user-api/internal/svc"
	"FirstProject/user-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func handlerGetUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHandlerGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.HandlerGetUserInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
