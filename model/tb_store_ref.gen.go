// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStoreRef = "tb_store_ref"

// StoreRef mapped from table <tb_store_ref>
type StoreRef struct {
	ID       int32  `gorm:"column:id;primaryKey" json:"id"`
	StoreKey string `gorm:"column:store_key" json:"store_key"`
	RefCount int32  `gorm:"column:ref_count" json:"ref_count"`
	Ctime    int32  `gorm:"column:ctime" json:"ctime"`
	Mtime    int32  `gorm:"column:mtime" json:"mtime"`
}

// TableName StoreRef's table name
func (*StoreRef) TableName() string {
	return TableNameStoreRef
}
