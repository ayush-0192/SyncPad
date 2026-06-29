package main 

import (
	"log"
	"github.com/ayush-0192/SyncPad/src/api/config"
	"github.com/ayush-0192/SyncPad/src/api/database"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ayush-0192/SyncPad/src/api/router"
	"github.com/ayush-0192/SyncPad/src/api/document"
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

	err = db.AutoMigrate(&document.Document{})	

	fmt.Printf("database connected successfully\n")

	
	_ = db
    
	docRepo := document.NewRepo(db)
	docService := document.NewService(docRepo)
	docHandler := document.NewHandler(docService)
	
	router := router.Setup(docHandler)
	router.Run(":" + cfg.PORT)
	
}