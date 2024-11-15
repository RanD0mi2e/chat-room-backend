package main

import (
	"chat-app-backend-service/internal/repository"
	"chat-app-backend-service/pkg/config"
	"chat-app-backend-service/pkg/log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("Server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	// 初始化MySQL
	db, err := sqlx.Connect("mysql", cfg.MySQL.DSN)
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	// 用户数据库
	userRepo := repository.NewUserRepo(db)
}
