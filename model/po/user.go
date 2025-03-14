package po

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Age      int    `json:"age,omitempty"`
}

func (User) TableName() string {
	return "user"
}
