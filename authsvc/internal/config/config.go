package config

import (
	"database/sql"
	"fmt"
	"os"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func (d *DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Pass, d.Name)
}

func (d *DBConfig) Configure(db *sql.DB) {
	// TODO: what are the optimum values?
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(5)
	// db.SetConnMaxLifetime(5)
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
		Port: ":8081",
		DB: &DBConfig{
			Host: "localhost",
			Port: "5432",
			User: "nexus",
			Pass: "nexus",
			Name: "nexus",
		},
		JWT: &JWTConfig{
			Secret: "secret",
		},
	}
}
