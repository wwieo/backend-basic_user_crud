package model

type User struct {
	ID       int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id"`
	Username string `gorm:"type:varchar(20) NOT NULL;" json:"username"`
	Password string `gorm:"type:varchar(100) NOT NULL;" json:"password"`
}
