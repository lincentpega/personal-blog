package ports

import (
	"content/post"
	"context"
)

type PostRepository interface {
	AddPost(context.Context, *post.Post) error
	UpdatePost(context.Context, *post.Post) error
	GetPostById(context.Context, string) (*post.Post, error)
	DeletePostById(context.Context, string) error
}
