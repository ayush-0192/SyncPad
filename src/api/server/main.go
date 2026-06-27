package main 

import (
	"log"
	"github.com/ayush-0192/SyncPad/src/api/config"
	"github.com/ayush-0192/SyncPad/src/api/database"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ayush-0192/SyncPad/src/api/router"
)

func main() {
	 err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	cfg := config.Load()

	fmt.Printf("%+v\n", cfg)

	db, err := database.ConnectPostgres(cfg)

	if err != nil {
		//fmt.printf("Error in connecting to database",err)
		log.Fatal("Error in connecting to database",err)
	}

	fmt.Printf("database connected successfully\n")

	
	_ = db
    
	
	router := router.Setup()
	router.Run(":" + cfg.PORT)
	
}