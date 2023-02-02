package main

import (
	"log"
	"strconv"

	"github.com/nlsh710599/go-practice/internal/config"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/route"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {

	rds, err := database.New(config.Get().PostgresHost, config.Get().PostgresUser, config.Get().PostgresPassword,
		config.Get().PostgresDatabase, config.Get().PostgresPort)

	if err != nil {
		log.Panicf("Failed to initialize RDS: %v", err)
	}

	if err := rds.CreateTable(); err != nil {
		log.Panicf("Failed to create table: %v", err)
	}

	controller := &route.Controller{
		RDS: rds,
	}

	r := gin.Default()

	route.Setup(r, controller)

	endless.ListenAndServe(":"+strconv.Itoa(config.Get().Port), r)
}
