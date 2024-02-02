package di

import (
	"github.com/anjush-bhargavan/microservice_gRPC_user/config"
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/db"
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/handler"
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/repo"
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/server"
	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/service"
)

// Init function load the configurations and initialize all instances
func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	userRepo := repo.NewUserRepo(db)
	userSvc := service.NewUserService(userRepo)
	userHandle := handler.NewUserHandler(userSvc)
	server.NewGrpcServer(userHandle)

}
