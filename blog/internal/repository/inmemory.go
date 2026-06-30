package repository

import (
	"easyapi/blog/internal/entity"
	"errors"
	"slices"
	"time"
)

var (
	ErrNoSuchPost = errors.New("no post with such id")
)

type inmemoryRepository struct {
	posts []entity.Post
}

func New() *inmemoryRepository {
	return &inmemoryRepository{
		posts: make([]entity.Post, 0),
	}
}

func (r *inmemoryRepository) GetPosts() ([]entity.Post, error) {
	return r.posts, nil
}

func (r *inmemoryRepository) CreatePost(post entity.Post) ([]entity.Post, error) {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	r.posts = append(r.posts, post)
	return r.posts, nil
}

func (r *inmemoryRepository) UpdatePostByID(id int, post entity.Post) ([]entity.Post, error) {
	if id >= len(r.posts) || id < 0 {
		return nil, ErrNoSuchPost
	}
	r.posts[id] = post
	return r.posts, nil
}

func (r *inmemoryRepository) GetPostByID(id int) (entity.Post, error) {
	if id >= len(r.posts) || id < 0 {
		return entity.Post{}, ErrNoSuchPost
	}

	return r.posts[id], nil
}

func (r *inmemoryRepository) DeletePostByID(id int) ([]entity.Post, error) {
	if id >= len(r.posts) || id < 0 {
		return nil, ErrNoSuchPost
	}

	r.posts = slices.Delete(r.posts, id, id+1)

	return r.posts, nil
}
