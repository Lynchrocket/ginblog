package models

import (
	"database/sql"
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassword,
		utils.DbName,
		utils.DbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("链接数据库失败，请检查参数：", err)
	}

	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		fmt.Printf("获取通用数据库对象失败，请检查参数：", err)
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//if err = sqlDB.Close(); err != nil {
	//	fmt.Printf("关闭数据库对象失败，请检查参数：", err)
	//}
}
