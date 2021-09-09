package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myGinWeb/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int           `gorm:"primaryKey" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

func InitDb(DbConfig setting.DatabaseConfig)(err error){
	db, err = gorm.Open(DbConfig.Type, fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DbConfig.Username,
		DbConfig.Host,
		DbConfig.DbName))
	if err != nil {
		return
	}
	db.AutoMigrate(&User{})
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db.DB().Ping()
}

func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedAt"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:created_at"); !ok {
		scope.SetColumn("UpdateAt", time.Now().Unix())
	}
}
