package main

import (
	"fmt"

	"github.com/prppoomw/blog-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
	dbClient := config.ConnectMongodb(cfg.MongoDBHost)
	fmt.Println(dbClient)
	defer config.CloseDatabaseConnection(dbClient)
}
