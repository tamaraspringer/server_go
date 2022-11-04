package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"  env-default:"true"`
	Listen  struct {
		Type   string `yaml:"type"  env-default:"port"`
		BindIp string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8081"`
	} `yaml: "listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("read config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config/config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatalln(err)

		}
	})
	return instance
}
