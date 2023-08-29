package config

import (
	"os"

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
		SSL       bool   `toml:"ssl"`
	} `toml:"s3"`
}

func Load(default_path string) (config Config, err error) {
	path := os.Getenv("CONFIG")

	if path == "" {
		path = default_path
	}

	_, err = toml.DecodeFile(path, &config)

	return
}
