package main

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/xgotyou/api_gateway/internal/http"
	"github.com/xgotyou/api_gateway/internal/services"
)

type Config struct {
	UserServiceAddr string `env:"USER_SERVICE_ADDR"`
	Port            int    `env:"PORT" env-default:"8081"`
	JwtSecret       string `env:"JWT_SECRET" env-required:""`
}

func main() {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
	}

	us := services.NewUserService(cfg.UserServiceAddr)

	r := http.SetupRouter(us)
	r.Use(http.JwtAuthHandler(cfg.JwtSecret))
	_ = r.Run(fmt.Sprintf(":%v", cfg.Port))
}
