package config

import (
	"flag"
	"fmt"
)

type DatabaseConfig struct {
	Host           string
	User           string
	Password       string
	DbName         string
	Ssl            bool
	ConnectTimeout uint
}

type ServerConfig struct {
	Address string
	Port    uint
	Cert    string
	Key     string
}

func (config *DatabaseConfig) ConnectionString() string {
	var sslmode string
	if config.Ssl {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		config.Host, config.User, config.Password, config.DbName, sslmode, config.ConnectTimeout)
}

func (config *ServerConfig) BindAddress() string {
	return fmt.Sprintf("%s:%d", config.Address, config.Port)
}

func GetDatabaseConfig() *DatabaseConfig {
	return &databaseConfig
}

func GetServerConfig() *ServerConfig {
	return &serverConfig
}

var databaseConfig = DatabaseConfig{}
var serverConfig = ServerConfig{}

func init() {
	flag.StringVar(&serverConfig.Address, "address", "localhost", "Bind host address")
	flag.UintVar(&serverConfig.Port, "port", 8080, "Listen port")
	flag.StringVar(&serverConfig.Cert, "cert", "", "SSL certificate file")
	flag.StringVar(&serverConfig.Key, "key", "", "Certificate key")

	flag.StringVar(&databaseConfig.Host, "db_host", "localhost", "Database host address")
	flag.StringVar(&databaseConfig.User, "db_user", "", "Database user name")
	flag.StringVar(&databaseConfig.Password, "db_password", "", "Database user password")
	flag.StringVar(&databaseConfig.DbName, "db_name", "", "Database name")
	flag.BoolVar(&databaseConfig.Ssl, "ssl", false, "SSL mode")
	flag.UintVar(&databaseConfig.ConnectTimeout, "connect_timeout", 2, "Database connect timeout")

	flag.Parse()
}
