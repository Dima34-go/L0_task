package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type DB struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	DBName   string `toml:"db_name"`
}
type ServerHTTP struct {
	Port string `toml:"port"`
}
type NATS struct {
	ClusterID   string `toml:"cluster_id"`
	ClientID    string `toml:"client_id"`
	Subject     string `toml:"subject"`
	QGroup      string `toml:"q_group"`
	DurableName string `toml:"durable_name"`
}
type Config struct {
	DB         `toml:"db"`
	ServerHTTP `toml:"server_http"`
	NATS       `toml:"nats"`
}
func InitConfig() (Config,error) {
	fContent, err := ioutil.ReadFile("configs/config.toml")
	if err != nil {
		return Config{},err
	}
	cfg:=Config{}
	_,err = toml.Decode(string(fContent),&cfg)
	return  cfg,err
}