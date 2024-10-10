package types

type Post struct {
	Id      int         `json:"id" gorm:"primary_key"`
	Title   interface{} `json:"title,omitempty" gorm:"default:'No title'"`
	Content string      `json:"content" gorm:"not null"`
}

func (p *Post) GetId() int {
	return p.Id
}

func (p *Post) GetTitle() interface{} {
	return p.Title
}

func (p *Post) GetContent() string {
	return p.Content
}
