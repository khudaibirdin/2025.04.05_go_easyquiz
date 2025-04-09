package config

import (
	"crypto/rsa"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTP struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Privatekey *rsa.PrivateKey
		PublicKey  *rsa.PublicKey
	} `yaml:"http"`
	Database struct {
		MySQL struct {
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
	} `yaml:"database"`
}

var (
	config *Config = &Config{}
)

func Get() *Config {
	return config
}

func New(path string) *Config {
	filename, err := filepath.Abs(path)
	if err != nil {
		panic(fmt.Sprintf("error with absolute path for path: %s, error: %s", path, err))
	}
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("error with file reading for path: %s, error: %s", filename, err))
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(fmt.Sprintf("error with yaml parsing error: %s", err))
	}
	privateKey, err := loadPrivateKey()
	if err != nil {
		panic(fmt.Sprintf("Ошибка загрузки ключа: %s", err))
	}
	config.HTTP.Privatekey = privateKey
	config.HTTP.PublicKey = &privateKey.PublicKey
	return config
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile("configs/private.pem")
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
}
