package types

type Post struct {
	Id      int         `json:"id" gorm:"primary_key"`
	Title   interface{} `json:"title,omitempty" gorm:"type:text;default:'No title'"`
	Content string      `json:"content" gorm:"type:text;not null"`
}

func (p *Post) TableName() string {
	return "posts"
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
