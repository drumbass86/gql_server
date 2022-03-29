package db

import (
	"fmt"
	"testing"

	"gorm.io/gorm/logger"
)

func TestInitDB(t *testing.T) {
	t.Log("Testing connection to psql db")
	err := InitDB()
	if err != nil {
		t.Error(err)
	}
}

func TestTableDB(t *testing.T) {
	t.Log("Testing tables exist")
	if LinksDB == nil {
		TestInitDB(t)
	}

	var tbCount int
	LinksDB.Raw("select COUNT(tablename) from pg_tables where tablename=?", User{}.TableName()).Scan(&tbCount)
	if tbCount != 1 {
		t.Errorf("Table '%s' dosn`t exist in DB!", User{}.TableName())
	}

	LinksDB.Raw("select COUNT(tablename) from pg_tables where tablename=?", Link{}.TableName()).Scan(&tbCount)
	if tbCount != 1 {
		t.Errorf("Table '%s' dosn`t exist in DB!", Link{}.TableName())
	}

	var resUsers User
	LinksDB.Raw("SELECT id, username, password from users WHERE id=?", 1).Scan(&resUsers)
	// Alternative select data from table
	// row := LinksDB.Raw("SELECT id, username, password from users WHERE id=?", 1).Row()
	// row.Scan(&resUsers.ID, &resUsers.Username, &resUsers.Password)
	if resUsers.ID <= 0 || resUsers.Username == "" {
		t.Error("Can`t select user data from DB table")
	}
	var resLink Link
	LinksDB.Raw("Select  * from links where userid=?", resUsers.ID).Scan(&resLink)
	if resLink.ID <= 0 {
		t.Errorf("Can`t select link data from DB table where userid=%v", resUsers.ID)
	}
	// Select full information about link
	LinksDB.Logger.LogMode(logger.Info)
	var lFull LinkFull
	if err := LinksDB.Preload("User_").Joins("left join users on links.userid=users.id", LinksDB.Where(&User{Username: "test"})).
		Find(&lFull).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", lFull)

}
