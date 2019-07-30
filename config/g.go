package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "go-mega"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString 返回数据库连接信息
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true", usr, pwd, host, port, db, charset)
}

// GetSMTPConfig SMTP配置信息
func GetSMTPConfig() (server string, port int, user, pwd string) {
	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

// GetServerURL
func GetServerURL() (url string) {
	url = viper.GetString("server.url")
	return
}
