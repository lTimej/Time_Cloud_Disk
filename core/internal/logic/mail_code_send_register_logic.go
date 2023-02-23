package logic

import (
	"context"
	"errors"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
	"liujun/Time_Cloud_Disk/core/models"
	"github.com/zeromicro/go-zero/core/logx"
	"liujun/Time_Cloud_Disk/core/helper"
	"time"
	"liujun/Time_Cloud_Disk/core/define"
	"log"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	count,err := l.svcCtx.DB.Where("email = ?",req.Email).Count(new(models.UserBasic))
	if err != nil{
		return
	}
	if count > 0{
		err = errors.New("邮箱已注册")
		return
	}
	code := helper.GetCode()
	l.svcCtx.RED.Set(l.ctx,req.Email,code,time.Second*time.Duration(define.CodeExpire))
	err = helper.SendEmail(req.Email,code)
	if err != nil{
		log.Println(err)
		err = errors.New("邮件发送失败") 
		return
	}
	resp.Code = 0
	resp.Info = ""
	return
}
