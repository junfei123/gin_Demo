package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"ginDemo/dao"
	"ginDemo/model"
	"time"
)

// 单例模式
type createUserService struct {
}

var CreateUserService = &createUserService{}

func (s *createUserService) CreateUser(ctx context.Context, username string, password string) error {

	//密码加密
	hash := sha256.Sum256([]byte(password))

	hashStr := hex.EncodeToString(hash[:])

	user := model.User{
		Name:     username,
		Password: hashStr,
		Ctime:    int32(time.Now().Unix()),
		Mtime:    int32(time.Now().Unix()),
	}

	return dao.Q.User.WithContext(ctx).Create(&user)

}
