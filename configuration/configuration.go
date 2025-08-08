package configuration

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Server struct {
	PORT string
}

type Configuration struct {
	SERVER Server
}

func Load(file string) (*Configuration, error) {
	err := godotenv.Load(file)
	if err != nil {
		fmt.Println(".env file not found. Operating system environment variables will be loaded")
	}
	return &Configuration{
		SERVER: Server{
			PORT: os.Getenv("SERVER_PORT"),
		},
	}, nil
}
