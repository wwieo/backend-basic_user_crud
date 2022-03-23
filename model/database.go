package model

type MysqlConfig struct {
	Url      string
	UserName string
	Password string
	Database string
	Table    string
	Port     int
}
