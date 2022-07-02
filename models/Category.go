package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Cid  uint8  `gorm:"type: primary_key; int; auto_increment" json:"cid"` // 分成第几类
	Name string `gorm:"type: varchar(20); not null" json:"name"`
}
