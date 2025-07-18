package service

import (
	"ginDemo/dao"
	"ginDemo/model"
	"time"

	"github.com/gin-gonic/gin"
)

type createFolderService struct {
}

var CreateFolderService = &createFolderService{}

func (s *createFolderService) CreateFolder(c *gin.Context, name string, parentId int) (*model.File, error) {
	folder := model.File{
		UserID:   int32(c.GetInt("user_id")),
		ParentID: int32(parentId),
		Name:     name,
		Type:     "folder",
		Ctime:    int32(time.Now().Unix()), // 创建时间
		Mtime:    int32(time.Now().Unix()), // 修改时间
	}
	if err := dao.Q.File.WithContext(c.Request.Context()).Create(&folder); err != nil {
		return nil, err
	}
	return &folder, nil
}
