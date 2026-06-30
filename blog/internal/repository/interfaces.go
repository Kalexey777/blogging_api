package repository

import (
	"context"
	"easyapi/blog/internal/entity"
)

type PostRepository interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error)
	UpdatePostByID(ctx context.Context, id int, post entity.Post) (*entity.Post, error)
	GetPostByID(ctx context.Context, id int) (*entity.Post, error)
	DeletePostByID(ctx context.Context, id int) (*entity.Post, error)
}
