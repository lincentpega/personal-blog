package post

import "time"

type Post struct {
	id        string
	text      string
	createdAt time.Time
	updatedAt time.Time
}

func New(id string, text string, createdAt time.Time, updatedAt time.Time) *Post {
	return &Post{id: id, text: text, createdAt: createdAt, updatedAt: updatedAt}
}

func NewWithText(text string) *Post {
	return &Post{
		text:      text,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}
}

func (p *Post) Id() string {
	return p.id
}

func (p *Post) Text() string {
	return p.text
}

func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Post) UpdatedAt() time.Time {
	return p.createdAt
}
