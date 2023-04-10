package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/xgotyou/api_gateway/internal/http"
	"github.com/xgotyou/api_gateway/internal/services"
)

type Config struct {
	UserServiceAddr string `env:"USER_SERVICE_ADDR"`
	Port            int    `env:"PORT" env-default:"8081"`
}

func main() {
	var cfg Config
	cleanenv.ReadEnv(&cfg)

	us := services.NewUserService(cfg.UserServiceAddr)

	r := http.SetupRouter(us)
	_ = r.Run(fmt.Sprintf(":%v", cfg.Port))
}
