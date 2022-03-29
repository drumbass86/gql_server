package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	hostDB                 = "172.17.0.2"
	userDB                 = "user"
	pwdDB                  = "123qwe"
	nameDB                 = "gql_db"
	portDB                 = 5432
	MAX_COUNT_LINKS_RETURN = 100
)

var LinksDB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		hostDB, userDB, pwdDB, nameDB, portDB)
	var err error
	LinksDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func CreateLink(l *Link) (uint, error) {
	if len(l.Address) > 0 && len(l.Title) > 0 && l.UserID > 0 {
		res := LinksDB.Create(l)
		return uint(res.RowsAffected), res.Error
	} else {
		return 0, errors.New("Incorrect parameter values in Link")
	}
}

func CreateUser(u *User) (uint, error) {
	if len(u.Password) > 0 && len(u.Username) > 0 {
		res := LinksDB.Create(u)
		return uint(res.RowsAffected), res.Error
	} else {
		return 0, errors.New("Incorrect parameter values in User")
	}
}

func GetAllLinks() ([]Link, error) {
	var allLinks []Link
	res := LinksDB.Model(&Link{}).Where(&Link{UserID: 1}).Limit(MAX_COUNT_LINKS_RETURN).Find(&allLinks)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return allLinks, nil
	}
}

func GetAllFullLinks() ([]LinkFull, error) {
	var allLinks []LinkFull
	res := LinksDB.Preload("User_").Joins("LEFT JOIN users on links.userid=users.id and users.id=?", 1).
		Limit(MAX_COUNT_LINKS_RETURN).Find(&allLinks)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return allLinks, nil
	}
}
