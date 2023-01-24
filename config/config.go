package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg *Config

type Config struct {
	LogLevel string `yaml:"log_level"`
	CaPath   string `yaml:"ca_path"`

	APIAdress string `yaml:"api_adress"`
	APIPort   int    `yaml:"api_port"`

	RedisHost          string `yaml:"redis_host"`
	RedisPassword      string `yaml:"redis_password"`
	RedisPort          string `yaml:"redis_port"`
	RedisLongLifetime  int    `yaml:"redis_long_lifetime"`
	RedisMidLifetime   int    `yaml:"redis_mid_lifetime"`
	RedisShortLifetime int    `yaml:"redis_short_lifetime"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalln(err)
	}
	Cfg = &t
}
