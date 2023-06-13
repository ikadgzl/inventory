package config

type Config struct {
	Port       string
	AuthSvcUrl string
	UserSvcUrl string
}

func NewConfig() *Config {
	// return &Config{
	// 	Port: os.Getenv("PORT"),
	// 	AuthSvcUrl: os.Getenv("AUTH_SVC_URL"),
	// 	UserSvcUrl: os.Getenv("USER_SVC_URL"),
	// }

	// TODO: for development.
	return &Config{
		Port:       ":8080",
		AuthSvcUrl: "localhost:8081",
		UserSvcUrl: "localhost:8082",
	}
}
