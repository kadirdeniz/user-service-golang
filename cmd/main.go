package main

import (
	"fmt"
	"user-service-golang/internal/handler"
	"user-service-golang/internal/repository"
	"user-service-golang/internal/service"
	"user-service-golang/pkg"
	"user-service-golang/tools"
)

func main() {
	fmt.Println("User Service")

	pkg.ReadConfigs()
	pkg.NewLogger()

	handler.NewAuthHandler(
		service.NewAuthService(
			service.NewUserService(repository.NewUserRepository(
				repository.NewUserDBRepository(tools.GetPostgresql()),
				repository.NewUserRedisRepository(tools.GetRedis()),
			)),
		),
	)

	tools.NewServer().Listen(":" + pkg.Configs.Application.Port)
}
