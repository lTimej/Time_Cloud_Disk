package logic

import (
	"context"
	"liujun/Time_Cloud_Disk/core/define"
	"liujun/Time_Cloud_Disk/core/models"

	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, user_identity string) (resp *types.UserFileListResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserFileListResponse)
	page := req.Page
	page_size := req.PageSize
	if page <= 0 {
		page = 1
	}
	if page_size == 0 {
		page_size = define.PageSize
	}
	offset := (page - 1) * page_size
	ur := models.UserRepository{}
	_, err = l.svcCtx.DB.Where("identity = ?", req.Identity).Select("id").Get(&ur)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	user_files := make([]*types.UserFile, 0)
	err = l.svcCtx.DB.Table(models.UserRepository{}).Where("user_repository.parent_id = ? AND user_repository.user_identity = ?", ur.Id, user_identity).Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,user_repository.name, repository_pool.path, repository_pool.size").Limit(page_size, offset).Find(user_files)
	if err != nil {
		resp.Info = "数据库错误"
		resp.Code = 1
		return resp, nil
	}
	resp.Count = len(user_files)
	resp.FileList = user_files
	resp.Code = 0
	return
}
