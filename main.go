package main

import (
	"dragondarkon/customers-api/apis"
	"dragondarkon/customers-api/database"
	"dragondarkon/customers-api/model"
	"dragondarkon/customers-api/router"
	"log"
)

func main() {
	db, err := database.Init("test.db")
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&model.Customer{})
	route := router.CustomersRouter{CustomersAPIs: apis.CustomerHandler{DB: db}}

	route.Route().Run(":" + "8080")
}
