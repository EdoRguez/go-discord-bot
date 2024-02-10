package config

import "github.com/spf13/viper"

type Config struct {
	TOKEN             string `mapstructure:"DISCORD_BOT_TOKEN"`
	DISCORD_SERVER_ID string `mapstructure:"DISCORD_SERVER_ID"`
	DEFAULT_CHANNEL   string `mapstructure:"DEFAULT_CHANNEL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
