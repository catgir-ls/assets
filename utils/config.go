package utils

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	App struct {
		Port int `toml:"port"`
	} `toml:"app"`

	S3 struct {
		Endpoint  string `toml:"endpoint"`
		Bucket    string `toml:"bucket"`
		AccessKey string `toml:"access_key"`
		SecretKey string `toml:"secret_key"`
	} `toml:"s3"`
}

var (
	cfg Config
)

func LoadConfig(path string) (err error) {
	_, err = toml.DecodeFile(path, &cfg)

	return
}

func GetConfig() Config {
	return cfg
}
