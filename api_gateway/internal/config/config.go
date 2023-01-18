package config

import (
	"fmt"
	"strings"

	"github.com/setest/pet-shop/api-gateway/internal/log"
	"github.com/setest/pet-shop/api-gateway/internal/resources"
	"github.com/spf13/viper"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	Logger    log.Config
	Resources resources.Config
	// TODO add list of services
	//AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	//ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	//OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	con := Config{}
	err = viper.Unmarshal(&con)
	if err != nil {
		panic(fmt.Errorf("cant unmarshal config: %w", err))
	}

	return &con, nil
}
