package dao

import (
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"fmt"
	"github.com/spf13/viper"
)

// var DB *gorm.DB
func Init() {
	GetConfig()
}
func GetConfig() {
	viper.SetConfigFile("dbConfig.json") //文件名
	viper.Set("Address", "0.0.0.0:9090") //统一把Key处理成小写 Address->address
	err := viper.WriteConfig()           //写入文件
	if err != nil {                      // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// func ConnectDB() {

// }
