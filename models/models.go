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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
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
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	return db.DB().Ping()
}

func CloseDB() {
	defer db.Close()
}