// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUser = "tb_user"

// User mapped from table <tb_user>
type User struct {
	ID       int32  `gorm:"column:id;primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Password string `gorm:"column:password" json:"password"`
	Ctime    int32  `gorm:"column:ctime" json:"ctime"`
	Mtime    int32  `gorm:"column:mtime" json:"mtime"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
