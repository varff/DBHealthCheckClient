package configs

import "github.com/joho/godotenv"

func LoadEnv() error {
	err := godotenv.Load("configs/config.env")
	if err != nil {
		return err
	}
	return nil
}
