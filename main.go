package main

import (
	"chat-app-backend-service/internal/handler"
	"chat-app-backend-service/internal/repository"
	"chat-app-backend-service/internal/server"
	"chat-app-backend-service/internal/service"
	"chat-app-backend-service/pkg/config"
	"chat-app-backend-service/pkg/log"
	"chat-app-backend-service/pkg/server/http"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func main() {
	// 配置层
	conf := config.NewConfig()

	// 日志层
	logger := log.NewLog(conf)
	logger.Info("Server start", zap.String("host", "http://localhost:"+conf.GetString("http.port")))
	// 数据库层
	mysqlDb, err := sqlx.Connect("mysql", conf.GetString("data.mysql.dsn"))
	if err != nil {
		logger.Fatal("连接数据库失败", zap.Error(err))
	}

	// db层
	repo := repository.NewRepository(logger, mysqlDb)

	// 事务层
	transaction := repository.NewTransation(repo)

	// 基础业务层
	serviceInst := service.NewService(transaction, logger)

	// 基础控制层
	handlerInst := handler.NewHandler(logger)

	// 用户服务层
	userService := service.NewUserService(serviceInst, repository.NewUserRepo(repo))
	// 用户控制层
	userHandler := handler.NewUserHandler(handlerInst, userService)

	// http服务层
	httpServer := server.NewServerHttp(logger, userHandler)

	// 启动服务
	http.Run(httpServer, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
