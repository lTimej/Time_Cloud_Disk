package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/models"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderListLogic {
	return &UserFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderListLogic) UserFolderList(req *types.UserFolderListRequest) (resp *types.UserFolderListResponse, err error) {
	// todo: add your logic here and delete this line
	ur := new(models.UserRepository)
	resp = new(types.UserFolderListResponse)
	ufl := []*types.UserFolder{}
	_, err = l.svcCtx.DB.Where("identity = ?", req.Identity).Select("id").Get(ur)
	if err != nil {
		resp.Code = 1
		resp.Info = "数据库错误"
		return resp, nil
	}
	err = l.svcCtx.DB.Where("parentId = ?", ur.Id).Select("identity,name").Find(ufl)
	if err != nil {
		resp.Code = 1
		resp.Info = "数据库错误"
		return resp, nil
	}
	resp.Code = 1
	resp.List = ufl
	return
}
