package db

type Link struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	UserID  uint   `json:"userid" gorm:"column:userid"`
}

func (l Link) TableName() string {
	return "links"
}

type LinkFull struct {
	Link
	User_ *User `gorm:"foreignKey:userid"`
}

func (l LinkFull) TableName() string {
	return "links"
}
