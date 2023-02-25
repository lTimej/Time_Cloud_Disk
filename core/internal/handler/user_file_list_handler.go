package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_Cloud_Disk/core/internal/logic"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
)

func UserFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err, 2222222)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Identity = r.URL.Query().Get("identity")
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.PageSize, _ = strconv.Atoi(r.URL.Query().Get("page_size"))
		l := logic.NewUserFileListLogic(r.Context(), svcCtx)
		resp, err := l.UserFileList(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
