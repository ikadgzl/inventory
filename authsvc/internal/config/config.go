package config

import "os"

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type JWTConfig struct {
	Secret    string
	Issuer    string
	ExpiresAt int64
}

type Config struct {
	Port string
	DB   *DBConfig
	JWT  *JWTConfig
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}

func NewConfig() *Config {
	// return &Config{
	// 	Port: os.Getenv("PORT"),
	// 	DB:   NewDBConfig(),
	// 	JWT:  NewJWTConfig(),
	// }

	// TODO: for dev.
	return &Config{
		Port: ":8080",
		DB: &DBConfig{
			Host: "localhost",
			Port: "5432",
			User: "test",
			Pass: "test",
			Name: "users",
		},
		JWT: &JWTConfig{
			Secret: "secret",
		},
	}
}
