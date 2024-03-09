package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	SecretKey string     `yaml:"SecretKey" env-required:"true"`
	Host      string     `yaml:"HOST" env-default:"127.0.0.1"`
	Port      string     `yaml:"PORT" env-default:"8081"`
	DBConfig  PostgresQL `yaml:"postgres" env-required:"true"`
	TLS       `yaml:"TLS" env-required:"true"`
}

type PostgresQL struct {
	Host       string  `yaml:"host" env-default:"127.0.0.1"`
	Port       string  `yaml:"port" env-default:"8080"`
	User       string  `yaml:"user" env-default:"postgres"`
	Password   string  `yaml:"password" env-default:"postgres"`
	DbName     string  `yaml:"db_name" env-default:"postgres"`
	Collection *string `yaml:"collection"`
}

type TLS struct {
	ServerKey string `yaml:"server_key" env-default:"./server.key"`
	ServerCRT string `yaml:"server_crt" env-default:"./server.crt"`
	Address   string `yaml:"address" env-default:":3000"`
	Network   string `yaml:"network" env-default:"tcp"`
}

func MustConfig() *Config {
	path := fetchConfigPath()

	if path == "" {
		path = "config/config.yaml"
	}

	if _, ok := os.Stat(path); os.IsNotExist(ok) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
