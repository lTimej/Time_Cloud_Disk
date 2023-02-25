package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/models"
	"log"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, user_identity string) (resp *types.UserFileDeleteResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserFileDeleteResponse)
	rp := new(models.RepositoryPool)
	ok, err := l.svcCtx.DB.Select("size").Where("identity = (select repository_identity from user_repository ur where ur.identity = ? limit 1)", req.Identity).Get(rp)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	if ok {
		if rp.Size > 0 {
			_, err = l.svcCtx.DB.Exec("update user_basic set now_volume = ? where identity = ?", rp.Size, user_identity)
			if err != nil {
				resp.Info = "数据库错误"
				resp.Code = 1
				log.Println(err)
				return resp, nil
			}
		}
	}
	_, err = l.svcCtx.DB.Where("identity = ? AND user_identity = ?", req.Identity, user_identity).Delete(&models.UserRepository{})
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	resp.Code = 1
	return
}
