package Config

import "os"

type Config struct {
	Port              string
	DatabaseURL       string
	PaymentGatewayURL string
}

func InitConfig() *Config {
	return &Config{
		Port:              os.Getenv("PORT"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		PaymentGatewayURL: os.Getenv("PAYMENT_GATEWAY"),
	}
}
