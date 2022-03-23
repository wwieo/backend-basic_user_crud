package dao

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "backend-user_crud/model"
)

var (
	Orm *gorm.DB
)

func GetConfig() (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath("./config")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Fatal("Config err: " + err.Error())
	}
	return config
}

func init() {
	mysqlConfig := *new(model.MysqlConfig)
	config := GetConfig()

	mysqlConfig = model.MysqlConfig{
		Url:      config.GetString("mysql.url"),
		UserName: config.GetString("mysql.username"),
		Password: config.GetString("mysql.password"),
		Database: config.GetString("mysql.database"),
		Port:     config.GetInt("mysql.port"),
	}

	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysqlConfig.UserName, mysqlConfig.Password,
		mysqlConfig.Url, mysqlConfig.Port, mysqlConfig.Database)

	var err error

	if Orm, err = gorm.Open(mysql.Open(addr), &gorm.Config{}); err != nil {
		log.Fatal("Connect mysql failed:", err)
		return
	}

	if _, err := Orm.DB(); err != nil {
		log.Fatal("Get mysql database failed:", err)
		return
	}

	Orm.Debug().AutoMigrate(&model.User{})
	migrator := Orm.Migrator()
	if !migrator.HasTable(&model.User{}) {
		fmt.Println("table not exist")
	}
}
