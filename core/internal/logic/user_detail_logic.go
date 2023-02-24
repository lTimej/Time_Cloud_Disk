package logic

import (
	"context"
	// "fmt"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
	"liujun/Time_Cloud_Disk/core/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	// todo: add your logic here and delete this line
	user_basic := models.UserBasic{}
	resp = new(types.UserDetailResponse)
	ok,err := l.svcCtx.DB.Where("identity = ?",req.Identity).Get(&user_basic)
	if err != nil{
		resp.Code = 1
		resp.Info = "数据库错误"
		return resp,nil
	}
	if !ok{
		resp.Code = 1
		resp.Info = "用户不存在"
		return resp,nil
	}
	resp.Code = 0
	resp.Name = user_basic.Name
	resp.Email = user_basic.Email
	return
}
