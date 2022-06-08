package mysql

import (
	"dousheng-backend/model"
	"dousheng-backend/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func Init(cfg *setting.MySQLConfig) error {
	//dsn := fmt.Sprintf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，
		},
	})

	if err != nil {
		fmt.Printf("数据库连接错误，请检查参数", err)
	}

	// 表迁移（创建表）
	if err := db.AutoMigrate(&model.Video{}, &model.User{}, &model.Comment{}, &model.RegistInfo{}); err != nil {
		fmt.Println("错误")
	}

	return nil
}
