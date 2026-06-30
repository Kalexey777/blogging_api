package controller

import (
	"context"
	"easyapi/blog/internal/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PostRepository interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
	CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error)
	UpdatePostByID(ctx context.Context, id int, post entity.Post) (*entity.Post, error)
	GetPostByID(ctx context.Context, id int) (*entity.Post, error)
	DeletePostByID(ctx context.Context, id int) (*entity.Post, error)
}

type controller struct {
	*gin.Engine
	logger *zap.Logger
	repo   PostRepository
}

func New(logger *zap.Logger, repo PostRepository) *controller {
	api := &controller{
		gin.Default(),
		logger,
		repo,
	}
	api.GET("/posts", api.getPosts)
	api.POST("/posts", api.createPost)
	api.PUT("/posts/:id", api.updatePostByID)
	api.GET("/posts/:id", api.getPostByID)
	api.DELETE("/posts/:id", api.deletePostByID)
	return api
}
