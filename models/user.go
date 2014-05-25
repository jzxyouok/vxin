package models

type User struct {
	Id            int64
	Username      string
	Password      string
	Email         string
	Role          int64
	Status        int64
	Remark        string
	LastLoginTime int64
	LastLoginIp   string
	LastLocation  string
}
