package model

type User struct {
	Username string `gorm:"type:varchar(20) NOT NULL;primary_key;" json:"username"`
	Password string `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
}
