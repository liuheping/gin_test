package context

import (
	"log"

	"github.com/spf13/viper"
)

type DbAuth struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	AppName    string
	AppVersion string

	DhbCommon DbAuth
	DhbData   DbAuth
	GinTest   DbAuth

	LogDebugMode bool
	LogFormat    string
}

func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("Config")
	config.AddConfigPath(path)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error context file: %s \n", err)
	}

	return &Config{
		AppName:    config.GetString("app_name"),
		AppVersion: config.GetString("app_version"),

		LogDebugMode: config.GetBool("log.log_debug_mode"),
		LogFormat:    config.GetString("log.log_format"),

		DhbCommon: DbAuth{
			Host:     config.GetString("dhb_common.host"),
			Port:     config.GetString("dhb_common.port"),
			User:     config.GetString("dhb_common.user"),
			Password: config.GetString("dhb_common.password"),
			DbName:   config.GetString("dhb_common.dbname"),
		},

		DhbData: DbAuth{
			Host:     config.GetString("dbb_data.host"),
			Port:     config.GetString("dbb_data.port"),
			User:     config.GetString("dbb_data.user"),
			Password: config.GetString("dbb_data.password"),
			DbName:   config.GetString("dbb_data.dbname"),
		},

		GinTest: DbAuth{
			Host:     config.GetString("db_gin_test.host"),
			Port:     config.GetString("db_gin_test.port"),
			User:     config.GetString("db_gin_test.user"),
			Password: config.GetString("db_gin_test.password"),
			DbName:   config.GetString("db_gin_test.dbname"),
		},
	}
}
