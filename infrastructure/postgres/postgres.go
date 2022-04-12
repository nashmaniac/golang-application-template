package postgres

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nashmaniac/golang-application-template/adapters"
	"github.com/nashmaniac/golang-application-template/models"
)

type postgresRepository struct {
	db *gorm.DB
}

func (r *postgresRepository) CloseDB(ctx context.Context) {
	sqlDb, _ := r.db.DB()
	log.Println("closing the database")
	sqlDb.Close()
}

func NewRepository(
	ctx context.Context,
	username string,
	password string,
	databaseName string,
	host string,
	port string,
) (adapters.PeristenceStore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, username, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("Connection to db is sucessful")
	if err != nil {
		return nil, err
	}

	// migration here
	db.AutoMigrate(&models.User{})

	return &postgresRepository{
		db: db,
	}, nil
}
