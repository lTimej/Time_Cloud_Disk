package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/helper"
	"liujun/Time_Cloud_Disk/core/models"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, user_identity string) (resp *types.UserRepositorySaveResponse, err error) {
	// todo: add your logic here and delete this line
	rp := new(models.RepositoryPool)
	resp = new(types.UserRepositorySaveResponse)
	_, err = l.svcCtx.DB.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	user_basic := new(models.UserBasic)
	_, err = l.svcCtx.DB.Select("now_volume,total_volume").Where("identity = ?", user_identity).Get(user_basic)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	if user_basic.NowVolume+rp.Size > user_basic.TotalVolume {
		resp.Info = "已超出总容量"
		resp.Code = 1
		return resp, nil
	}
	sql := "update user_basic set now_volume = ? where id = ?"
	_, err = l.svcCtx.DB.Exec(sql, user_basic.NowVolume+rp.Size, user_basic.Id)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	ur := models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       user_identity,
		RepositoryIdentity: rp.Identity,
		ParentId:           req.ParentId,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.DB.Insert(&ur)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	resp.Code = 0
	return
}
