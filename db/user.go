package db

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (m User) TableName() string {
	return "users"
}
