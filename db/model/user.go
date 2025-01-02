package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
}

func (User) TableName() string {
	return "user"
}
