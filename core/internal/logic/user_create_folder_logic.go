package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/helper"
	"liujun/Time_Cloud_Disk/core/models"
	"log"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateFolderLogic {
	return &UserCreateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateFolderLogic) UserCreateFolder(req *types.UserFolderCreateRequest, user_identity string) (resp *types.UserFolderCreateResponse, err error) {
	// todo: add your logic here and delete this line
	ur := new(models.UserRepository)
	resp = new(types.UserFolderCreateResponse)
	cn, err := l.svcCtx.DB.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(ur)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	if cn > 0 {
		resp.Info = "文件夹已存在j"
		resp.Code = 1
		return resp, nil
	}
	ur.Identity = helper.UUID()
	ur.UserIdentity = user_identity
	ur.Name = req.Name
	ur.ParentId = req.ParentId
	_, err = l.svcCtx.DB.Insert(ur)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		log.Println(err)
		return resp, nil
	}
	resp.Code = 0
	return
}
