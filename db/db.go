package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	hostDB = "172.17.0.2"
	userDB = "user"
	pwdDB  = "123qwe"
	nameDB = "gql_db"
	portDB = 5432
)

var LinksDB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostDB, userDB, pwdDB, nameDB, portDB)
	var err error
	LinksDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
