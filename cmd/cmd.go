package cmd

import (
	"backend-api/client/ws"
	"backend-api/config"
	"backend-api/db"
	"backend-api/handler"
	"backend-api/utils/logger"
	"fmt"
	"log"

	dbModel "backend-api/model/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	if cfg == nil {
		panic("empty config")
	}
	fmt.Println("✅ Config loaded")

	// logger
	logger.InitLogger(cfg.Logger.Path)
	fmt.Println("✅ Initiated logger")

	// db
	dbConn, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	fmt.Println("✅ Connected to DB")

	// Inject DB into domain
	db.InjectDB(dbConn)

	if err = dbModel.AutoMigrateAll(dbConn); err != nil {
		panic(err)
	}
	fmt.Println("✅ Migrated models")

	// router
	r := handler.SetupRouter(&cfg.Auth)
	log.Println("🚀 Server running on :", cfg.Port)

	// client
	ws.InitGRPCClient(&cfg.GrpcClient)
	fmt.Println("✅ Initiated grpc client")

	// run server
	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
