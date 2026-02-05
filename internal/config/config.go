package config

import (
	"time"
)

type Config struct {
	Env 		 string `yaml:"env" env-default:"local"`
    StoragePath  string `yaml:"storage_path" env-required:"true"`
	HttpServer   `yaml:"http_server"`
}

type HttpServer struct {
	Address 	 string `yaml:"address" env-default:"localhost:8283"`
    Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
    IdleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")
}
