package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type: varchar(100); not null" json:"title"` // 标题
	Desc     string   `gorm:"type: varchar(200)" json:"desc"`            // 描述
	Content  string   `gorm:"type: text" json:"content"`                 // 内容
	Img      string   `gorm:"type: varchar(100)" json:"img"`             // 图片
}
