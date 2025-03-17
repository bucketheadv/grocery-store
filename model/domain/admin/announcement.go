package admin

type Announcement struct {
	ID         int64  `json:"id"`
	Content    string `json:"content"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (p *Announcement) TableName() string {
	return "announcement_info"
}

func (p *Announcement) GetID() int64 {
	return p.ID
}
