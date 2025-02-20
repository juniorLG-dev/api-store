package main

import (
	"loja/internal/configuration/database"
	"loja/internal/configuration/cache_config"
	"loja/internal/configuration/initializer_http"

	"github.com/gin-gonic/gin"
	
	"log"
)

func main() {
	router := gin.Default()

	db, err := database.SetupDB()
	if err != nil {
		log.Panic(err)
	}

	rdb := cache_config.SetupCacheDB("localhost:6379", "", 0)

	initializer_http.InitSeller(db, &rdb, &router.RouterGroup)
	initializer_http.InitInventory(db, &router.RouterGroup)

	router.Run(":8080")
}