package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func (User) TableName() string {
	return "user"
}
