package di

import (
	server "user/service/pkg/api"
	"user/service/pkg/api/service"
	"user/service/pkg/config"
	"user/service/pkg/db"
	"user/service/pkg/repository"
	"user/service/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(adminRepository)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
