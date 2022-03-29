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

type User struct {
	ID       uint
	Username string
	Password string
}

func (m User) TableName() string {
	return "users"
}

type Link struct {
	ID      uint
	Title   string
	Address string
	UserID  string `gorm:"column:userid"`
}

func (l Link) TableName() string {
	return "links"
}

type LinkFull struct {
	ID      uint
	Title   string
	Address string
	UserID  string `gorm:"column:userid"`
	User_   User   `gorm:"foreignKey:userid"`
}

func (l LinkFull) TableName() string {
	return "links"
}

var LinksDB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostDB, userDB, pwdDB, nameDB, portDB)
	var err error
	LinksDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
