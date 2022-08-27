package config

import (
	"github.com/spf13/viper"
)

func Mode() string {
	return viper.GetString("mode")
}

func HttpDomain() string {
	return viper.GetString("httpDomain")
}

func HttpPort() string {
	return viper.GetString("httpPort")
}

func CouchDBPort() string {
	return viper.GetString("couchDBPort")
}

func DbUrl() string {
	return viper.GetString("dbUrl")
}

func DbPort() string {
	return viper.GetString("dbPort")
}

func DbUser() string {
	return viper.GetString("dbUser")
}

func DbName() string {
	return viper.GetString("dbName")
}

func DbPassword() string {
	return viper.GetString("dbPassword")
}

func SecretKey() string {
	return viper.GetString("secretKey")
}

func LoginExpTime() int64 {
	return viper.GetInt64("loginExpTime")
}

func LoginRefreshExpTime() int64 {
	return viper.GetInt64("loginRefreshExpTime")
}