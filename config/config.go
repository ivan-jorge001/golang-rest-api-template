package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type PostgressConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type Config struct {
	Port        int
	Environment string
	ApiKey      string
	JwtSecret   string
	Postgress   *PostgressConfig
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Default to development if no environment is specified
	}

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// Set up to automatically read configuration from environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	postgress := &PostgressConfig{}
	if err := viper.Sub("postgress").Unmarshal(postgress); err != nil {
		log.Fatalf("Error unmarshalling postgress config: %s", err)
	}

	return &Config{
		Port:        viper.GetInt("port"),
		Environment: viper.GetString("environment"),
		ApiKey:      viper.GetString("api_key"),
		JwtSecret:   viper.GetString("jwt_secret"),
		Postgress:   postgress,
	}
}

func (cfg *Config) IsDev() bool {
	return cfg.Environment == "development"
}
