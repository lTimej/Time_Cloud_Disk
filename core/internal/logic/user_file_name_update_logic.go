package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/models"
	"log"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, user_identity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	ur := new(models.UserRepository)
	resp = new(types.UserFileNameUpdateResponse)
	cn, err := l.svcCtx.DB.Where("name = ? AND parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Name, req.Identity).Count(ur)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	if cn > 0 {
		resp.Info = "文件名已存在"
		resp.Code = 1
		return resp, nil
	}
	data := &models.UserRepository{Name: req.Name}
	_, err = l.svcCtx.DB.Where("user_identity = ? AND identity = ?", user_identity, req.Identity).Update(data)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	resp.Code = 0
	return
}
