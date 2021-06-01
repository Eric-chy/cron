package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u User) TableName() string {
	return "user"
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var list []*User
	if u.Name != "" {
		db = db.Where("name = ?", u.Name)
	}
	err := db.Limit(pageSize).Offset(pageOffset).Find(&list).Error
	return list, err
}
