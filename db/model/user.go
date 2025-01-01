package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
}

func (User) TableName() string {
	return "user"
}
