package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(host, username, password, databaseName string, port int) (RDS, error) {
	instance, err := initializePostgres(host, username, password, databaseName, port)

	if err != nil {
		log.Printf("Failed to initialize rds: %v", err)
		return nil, err
	}

	return instance, nil
}

func initializePostgres(host, username, password, databaseName string, port int) (*postgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, username, password, databaseName, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to postgres: %v", err)
		return nil, err
	}

	return &postgresClient{client: db}, nil
}
