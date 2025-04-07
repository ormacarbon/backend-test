package config

import "github.com/spf13/viper"

type Config struct {
	PostgresHost  string `mapstructure:"POSTGRES_HOST"`
	PostgresPort  string `mapstructure:"POSTGRES_PORT"`
	PostgresUser  string `mapstructure:"POSTGRES_USER"`
	PostgresPass  string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb    string `mapstructure:"POSTGRES_DATABASE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	SMTPHost      string `mapstructure:"SMTP_HOST"`
	SMTPPort      int    `mapstructure:"SMTP_PORT"`
	SMTPUser      string `mapstructure:"SMTP_USER"`
	SMTPPass      string `mapstructure:"SMTP_PASS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
