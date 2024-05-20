package main

import (
	"fmt"
	"log"

	"github.com/wxw9868/video/initialize/db"
	"github.com/wxw9868/video/router"
	"github.com/wxw9868/video/utils"
)

func main() {
	db.RegisterTables()

	router := router.Engine()

	ip, err := utils.GetLocalIP()
	if err != nil {
		log.Fatal(err)
	}
	// ip = "127.0.0.1"

	if err := router.Run(fmt.Sprintf("%s:%d", ip, 80)); err != nil {
		panic(err)
	}
}
