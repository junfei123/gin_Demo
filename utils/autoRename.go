package utils

import (
	"context"
	"fmt"
	"ginDemo/dao"
	"strings"
)

func AutoRenameFile(c context.Context, userId int32, fileName string, parentId int32) string {
	name := fileName
	idx := 1
	ext := ""

	if dot := strings.LastIndex(fileName, "."); dot > -1 {
		ext = fileName[dot:]
		name = fileName[:dot]
	}

	for {
		count, err := dao.Q.File.WithContext(c).Where(dao.File.Name.Eq(fileName), dao.File.ParentID.Eq(parentId), dao.File.UserID.Eq(userId)).Count()
		if err != nil {
			fmt.Println("读取数据库错误：", err)
		}
		if count == 0 {
			return name + ext
		}
		fileName = fmt.Sprintf("%s(%d)%s", name, idx, ext)
		idx++
	}
}
