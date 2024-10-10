package configs

import (
	"errors"
	"os"
)

const (
	defaultApiPort       = "8080"
	defaultDbHost        = "localhost"
	defaultDbPort        = "5432"
	defaultDbUsername    = "root"
	defaultDbPassword    = "root"
	defaultDbDatabase    = "posts"
	defaultMongoProtocol = "mongodb://"
	defaultMongoHost     = "localhost"
	defaultMongoPort     = "27017"
)

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type MongoDBConfig struct {
	Protocol string
	Host     string
	Port     string
}

type Config struct {
	ApiConfig     APIConfig
	DbConfig      DBConfig
	MongoDbConfig MongoDBConfig
}

func getMongoConfig() *MongoDBConfig {
	mongoProtocol, ok := os.LookupEnv("MONGO_PROTOCOL")
	if !ok {
		mongoProtocol = defaultMongoProtocol
	}
	mongoHost, ok := os.LookupEnv("MONGO_HOST")
	if !ok {
		mongoHost = defaultMongoHost
	}
	mongoPort, ok := os.LookupEnv("MONGO_PORT")
	if !ok {
		mongoPort = defaultMongoPort
	}
	return &MongoDBConfig{
		Protocol: mongoProtocol,
		Host:     mongoHost,
		Port:     mongoPort,
	}
}

func getDatabaseConfig() *DBConfig {
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = defaultDbHost
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = defaultDbPort
	}

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		dbUsername = defaultDbUsername
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		dbPassword = defaultDbPassword
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbName = defaultDbDatabase
	}

	return &DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUsername,
		Password: dbPassword,
		Database: dbName,
	}
}

func getApiConfig() *APIConfig {
	apiPort, ok := os.LookupEnv("API_PORT")
	if !ok {
		apiPort = defaultApiPort
	}
	return &APIConfig{
		Port: apiPort,
	}
}

func LoadConfig() (*Config, error) {
	apiConfig := getApiConfig()
	dbConfig := getDatabaseConfig()
	mongoConfig := getMongoConfig()

	appConfig := &Config{
		ApiConfig:     *apiConfig,
		DbConfig:      *dbConfig,
		MongoDbConfig: *mongoConfig,
	}

	if appConfig != nil {
		return appConfig, nil
	}
	return nil, errors.New("app configuration not founded")
}

// to be used When optimized
// func getEnv(key, defaultValue string) string {
//     val, ok := os.LookupEnv(key); ok {
//     		return val
//     }
//     os.Setenv(key, defaultValue)
//     return defaultValue
// }
//
// func LoadConfig() *Config {
//     apiConfig := &APIConfig{
//         Port: getEnv("API_PORT", "8080"),
//     }
//
//     dbConfig := &DBConfig{
//         Host:     getEnv("DB_HOST", "localhost"),
//         Port:     getEnv("DB_PORT", "5432"),
//         Username: getEnv("DB_USER", "root"),
//         Password: getEnv("DB_PASSWORD", "root"),
//         Database: getEnv("DB_NAME", "posts"),
//     }
//
//     return &Config{
//         ApiConfig: *apiConfig,
//         DbConfig:  *dbConfig,
//     }
// }
//
