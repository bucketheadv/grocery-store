package domain

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Age        int    `json:"age,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
}

func (User) TableName() string {
	return "user_info"
}

func (u User) GetID() int64 {
	return u.ID
}
