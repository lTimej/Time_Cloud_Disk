package handler

import (
	"net/http"
	"crypto/md5"
	"fmt"
	"path"
	"log"
	"github.com/zeromicro/go-zero/rest/httpx"
	"liujun/Time_Cloud_Disk/core/internal/logic"
	"liujun/Time_Cloud_Disk/core/internal/svc"
	"liujun/Time_Cloud_Disk/core/internal/types"
	"liujun/Time_Cloud_Disk/core/models"
	"liujun/Time_Cloud_Disk/core/helper"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Println(11111,err)
			httpx.Error(w, err)
			return
		}
		b := make([]byte,fileHeader.Size)
		_,err = file.Read(b)
		if err != nil {
			log.Println(22222,err)
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x",md5.Sum(b))
		repository_pool := new(models.RepositoryPool)
		ok,err := svcCtx.DB.Where("hash = ?",hash).Get(repository_pool)
		if err != nil {
			log.Println(33333,err)
			httpx.Error(w, err)
			return
		}
		if ok {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: repository_pool.Identity, Ext: repository_pool.Ext, Name: repository_pool.Name})
			return
		}
		file_path,err := helper.MinIOUpload(r)
		if err != nil {
			log.Println(fileHeader.Filename,4444444,err)
			httpx.Error(w, err)
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = file_path
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
