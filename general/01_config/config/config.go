package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Config struct {
	Port     int         `json:"port"`
	Env      string      `json:"env"`
	Database MysqlConfig `json:"database"`
}

func LoadConfig() Config {
	f, err := os.Open("my_config.json")
	if err != nil {
		panic(err)
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}

func (c MysqlConfig) Dialect() string {
	return "mysql"
}

func (c MysqlConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("%s@/%s?parseTime=true", c.User, c.Name)
	}
	return fmt.Sprintf("%s:%s@/%s?parseTime=true", c.User, c.Password, c.Name)
}
