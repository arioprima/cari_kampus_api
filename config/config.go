package config

import (
	"log"

	"github.com/spf13/viper"
)

type Database struct {
	DBHost string `mapstructure:"host"`
	DBPort string `mapstructure:"port"`
	DBName string `mapstructure:"name"`
	DBUser string `mapstructure:"user"`
	DBPass string `mapstructure:"password"`
}

type Email struct {
	EmailFrom string `mapstructure:"from"`
	SMTPHost string `mapstructure:"host"`
	SMTPPass string `mapstructure:"password"`
	SMTPPort int `mapstructure:"port"`
	SMTPUser string `mapstructure:"user"`
}

type Config struct {
	PORT string `mapstructure:"PORT"`
	ClientOrigin string `mapstructure:"origin"`
	TokenSecret string `mapstructure:"secret"`
	TokenMaxAge int `mapstructure:"maxage"`
	Database Database `mapstructure:"database"`
	Email Email `mapstructure:"email"`
}

func LoadConfig(path string) (config Config, err error) {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err = viper.ReadInConfig();
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return config, nil
}