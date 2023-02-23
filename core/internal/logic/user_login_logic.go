package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/helper"
	"liujun/Time_Cloud_Disk/core/models"
	"time"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	// todo: add your logic here and delete this line
	username := req.Name
	password := req.Password
	user := new(models.UserBasic)
	resp = new(types.UserLoginResponse)
	ok, err := l.svcCtx.DB.Where("name= ? AND password = ?", username, helper.MD5(password)).Get(user)
	if err != nil {
		resp.Code = 1
		resp.Info = "数据库出错"
		return resp, nil
	}
	if !ok {
		resp.Code = 1
		resp.Info = "用户名或密码错误"
		return resp, nil
	}
	token, err := helper.GenToken(user.Id, user.Identity, user.Name, time.Hour*2)
	if err != nil {
		resp.Code = 1
		resp.Info = err.Error()
		return resp, nil
	}
	refresh_token, err := helper.GenToken(user.Id, user.Identity, user.Name, time.Hour*24*7)
	if err != nil {
		resp.Code = 1
		resp.Info = err.Error()
		return resp, nil
	}
	resp.Code = 0
	resp.Token = token
	resp.RefreshToken = refresh_token
	return
}
