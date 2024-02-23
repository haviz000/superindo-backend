package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/superindo-retail/database"
	"github.com/haviz000/superindo-retail/route"
	"github.com/joho/godotenv"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("load env failed")
	}

	database.ConnectDB()
	db := database.GetDB()

	route.Routes(router, db)

	router.Run(PORT)
}
