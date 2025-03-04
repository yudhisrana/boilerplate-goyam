package config

import (
	"os"

	"github.com/stretchr/testify/assert/yaml"
	handlerErrors "github.com/yudhisrana/boilerplate-goyam/internal/errors"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Port      string `yaml:"port"`
	SecretKey string `yaml:"secret_key"`
	Mode      string `yaml:"mode"`
}

type DBConfig struct {
	Host         string       `yaml:"host"`
	Port         string       `yaml:"port"`
	User         string       `yaml:"user"`
	Password     string       `yaml:"password"`
	Database     string       `yaml:"database"`
	DBConnection DBConnection `yaml:"db_connection"`
}

type DBConnection struct {
	MaxIdleConnections int `yaml:"max_idle_connections"`
	MaxOpenConnections int `yaml:"max_open_connections"`
	ConnMaxLifetime    int `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime    int `yaml:"conn_max_idle_time"`
}

var Cfg Config

func LoadConfig(path string) (err error) {
	configByte, err := os.ReadFile(path)
	if err != nil {
		return handlerErrors.ErrReadFileConfig
	}

	err = yaml.Unmarshal(configByte, &Cfg)
	if err != nil {
		return handlerErrors.ErrUnmarshalConfig
	}

	return nil
}
