package helpers

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}

func GenerateUUID() (string, error) {
	var uid string
	id, err := uuid.NewV7()
	if err != nil {
		return uid, err
	}

	return id.String()[:9], nil
}
