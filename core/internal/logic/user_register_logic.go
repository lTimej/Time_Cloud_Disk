package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"liujun/Time_Cloud_Disk/core/helper"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
	"liujun/Time_Cloud_Disk/core/models"
	"log"
	"strings"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	log.Println(req)
	resp = new(types.Response)
	redis_code, err := l.svcCtx.RED.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("该邮箱的验证码为空")
	}
	if strings.ToLower(redis_code) != strings.ToLower(req.Code) {
		return nil, errors.New("验证码错误")
	}
	count, err := l.svcCtx.DB.Where("name = ?", req.Name).Count(&models.UserBasic{})
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("用户已存在")
	}
	user_basic := models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.MD5(req.Password),
		Email:    req.Email,
	}
	n, err := l.svcCtx.DB.Insert(user_basic)
	if err != nil {
		return nil, err
	}
	log.Println(n)
	resp.Info = ""
	resp.Code = 0
	return
}
