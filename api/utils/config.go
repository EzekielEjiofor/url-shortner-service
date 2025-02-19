package utils

import "github.com/spf13/viper"

type Config struct {
	DBAddress string `mapstructure:"DB_ADDR"`
	Domain    string `mapstructure:"DOMAIN"`
	AppPort   string `mapstructure:"APP_PORT"`
	Quota     int    `mapstructure:"API_QUOTA"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
