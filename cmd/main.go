package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prppoomw/blog-api/internal/config"
	"github.com/prppoomw/blog-api/internal/route"
)

func main() {
	cfg := config.LoadConfig()
	dbClient := config.ConnectMongodb(cfg.MongoDBHost)
	db := dbClient.Database(cfg.DBName)
	defer config.CloseDatabaseConnection(dbClient)

	timeout := time.Duration(cfg.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(cfg, timeout, db, gin)

	gin.Run(cfg.ServerPort)
}
