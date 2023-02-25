package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_Cloud_Disk/core/internal/logic"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
)

func UserCreateFolderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserCreateFolderLogic(r.Context(), svcCtx)
		resp, err := l.UserCreateFolder(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
