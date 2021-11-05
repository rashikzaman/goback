package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configInstance *Config

type Config struct {
}

func GetConfig() (config Config) {
	if configInstance == nil {
		envFile := os.Getenv("ENV_FILE")
		if envFile == "" {
			envFile = ".env"
		}
		err := godotenv.Load(envFile)
		if err != nil {
			panic("Error loading .env file")
		}
		configInstance = new(Config)
	}
	return *configInstance
}

type DatabaseConfig struct {
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

func (Config) getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c Config) getBool(key string, fallback bool) bool {
	val := c.getEnv(key, strconv.FormatBool(fallback))
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		//logger.GetLog().Warn(fmt.Sprintf("config conversion to bool failed, key: %v", key))
		log.Fatal("config conversion to bool failed")
	}
	return boolVal
}

func (c Config) GetEnvVariable(key string) string {
	return c.getEnv(key, "")
}

func (c Config) GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DatabaseName: c.getEnv("MYSQL_DB_NAME", "locally"),
		Host:         c.getEnv("MYSQL_HOST", "localhost"),
		Port:         c.getEnv("MYSQL_PORT", "5432"),
		Username:     c.getEnv("MYSQL_DB_USERNAME", "locally"),
		Password:     c.getEnv("MYSQL_DB_PASSWORD", "locally"),
	}
}
