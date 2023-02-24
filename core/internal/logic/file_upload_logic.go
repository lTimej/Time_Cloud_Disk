package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/models"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
	"liujun/Time_Cloud_Disk/core/helper"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.FileUploadResponse)
	rp := &models.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	_,err = l.svcCtx.DB.Insert(rp)
	if err != nil{
		resp.Code = 1
		resp.Info = "数据库错误"
		return resp,nil
	}
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	resp.Code = 0
	return
}
