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
	Database     DatabaseConfig
	Port      int
}
type DatabaseConfig struct {
	Type string
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
	if key = os.Getenv("DATABASE_DB_NAME"); key != "" {
		Config.Database.Type = key
	}
	if key = os.Getenv("DATABASE_PORT"); key != "" {
		if b, err := strconv.Atoi(key); err == nil {
			Config.Database.Port = b
		}
	}
	if key = os.Getenv("DATABASE_HOST"); key != "" {
		Config.Database.Host = key
	}
	if key = os.Getenv("DATABASE_PASSWORD"); key != "" {
		Config.Database.Password = key
	}
	if key = os.Getenv("DATABASE_DB_NAME"); key != "" {
		Config.Database.DbName = key
	}
}
