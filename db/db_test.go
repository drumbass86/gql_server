package db

import "testing"

func TestInitDB(t *testing.T) {
	t.Log("Testing connection to psql db")
	err := InitDB()
	if err != nil {
		t.Error(err)
	}
}

type ResUser struct {
	id   uint   `gorm:"column:id"`
	name string `gorm:"column:username"`
	pwd  string `gorm:"column:password"`
}

func (m ResUser) TableName() string {
	return "users"
}

func TestTableDB(t *testing.T) {
	t.Log("Testing tables exist")
	if LinksDB == nil {
		TestInitDB(t)
	}

	var tbCount int
	LinksDB.Raw("select COUNT(tablename) from pg_tables where tablename='users'").Scan(&tbCount)
	if tbCount != 1 {
		t.Error("Table 'users' dosn`t exist in DB!")
	}

	LinksDB.Raw("select COUNT(tablename) from pg_tables where tablename='links'").Scan(&tbCount)
	if tbCount != 1 {
		t.Error("Table 'links' dosn`t exist in DB!")
	}

	var resUsers ResUser
	//Don`t work
	//LinksDB.Raw("SELECT id, username, password from users WHERE id=?", 1).Scan(&resUsers)
	row := LinksDB.Raw("SELECT id, username, password from users WHERE id=?", 1).Row()
	row.Scan(&resUsers.id, &resUsers.name, &resUsers.pwd)
	if resUsers.id <= 0 || resUsers.name == "" {
		t.Error("Can`t select users from DB table")
	}
}
