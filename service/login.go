package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"ginDemo/dao"
	"log"
)

// 单例
type loginService struct {
}

var LoginService = &loginService{}

// 登录成功返回userid
func (s *loginService) Login(ctx context.Context, username, password string) (int32, error) {

	user, err := dao.Q.User.WithContext(ctx).Where(dao.Q.User.Name.Eq(username)).First()

	if err != nil || user == nil {
		log.Println("用户不存在")
		return 0, errors.New("用户不存在")
	}

	hash := sha256.Sum256([]byte(password))

	hashStr := hex.EncodeToString(hash[:])
	if hashStr != user.Password {
		log.Println("密码错误")
		return 0, errors.New("密码错误")
	}

	return user.ID, nil
}
