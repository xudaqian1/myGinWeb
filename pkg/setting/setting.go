package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type IConfig struct {
	JwtSecret string
	RunMode   string
	MySQL     MySqlConfig
	Port      int
}
type MySqlConfig struct {
	Username string
	Password string
	DbName   string
	Host     string
	Port     int
}

var (
	Config IConfig
)

func init() {
	ViperConfig := viper.New()
	ViperConfig.AddConfigPath("./conf/")
	ViperConfig.SetConfigName("config")
	ViperConfig.SetConfigType("json")
	err := ViperConfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := ViperConfig.Unmarshal(&Config); err != nil {
		fmt.Println(err)
	}

	var key string
	if key = os.Getenv("RUN_MODE"); key != "" {
		Config.RunMode = key
	}
	if key = os.Getenv("PORT"); key != "" {
		if b, err := strconv.Atoi(key); err == nil {
			Config.Port = b
		}
	}
	if key = os.Getenv("JWT_SECRET"); key != "" {
		Config.JwtSecret = key
	}
	//mysql
	if key = os.Getenv("MYSQL_PORT"); key != "" {
		if b, err := strconv.Atoi(key); err == nil {
			Config.MySQL.Port = b
		}
	}
	if key = os.Getenv("MYSQL_HOST"); key != "" {
		Config.MySQL.Host = key
	}
	if key = os.Getenv("MYSQL_PASSWORD"); key != "" {
		Config.MySQL.Password = key
	}
	if key = os.Getenv("MYSQL_DB_NAME"); key != "" {
		Config.MySQL.DbName = key
	}
}
