package service

import (
	"ginDemo/dao"
	"ginDemo/model"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type uploadFileService struct {
}

var UploadFileService = &uploadFileService{}

const storagePath = "./data"

func (s *uploadFileService) UploadFile(c *gin.Context, file *multipart.FileHeader, fileName string, parentId int) (*model.File, error) {

	//1.打开上传的文件
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(storagePath, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	if _, err = io.Copy(out, f); err != nil {
		return nil, err
	}

	userId := c.GetInt("user_Id")

	fileRecord := model.File{
		Name:     fileName,
		ParentID: int32(parentId),
		UserID:   int32(userId),
		Size:     int32(file.Size),
		Type:     "file",
		StoreKey: filePath,
		VerNum:   1,
		Ctime:    int32(time.Now().Unix()),
		Mtime:    int32(time.Now().Unix()),
	}
	if err := dao.Q.File.WithContext(c).Create(&fileRecord); err != nil {
		os.Remove(filePath)
		return nil, err
	}

	return &fileRecord, nil
}
