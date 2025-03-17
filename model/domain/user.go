package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (u User) GetID() int {
	return u.ID
}
