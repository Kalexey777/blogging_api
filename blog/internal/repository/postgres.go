package repository

import (
	"context"
	"easyapi/blog/internal/entity"
	errorspkg "easyapi/blog/internal/errors"
	sqlc "easyapi/blog/internal/repository/sqlc"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type postgresRepository struct {
	queries sqlc.Querier
	logger  *zap.Logger
}

func NewPostgresRepository(qdb sqlc.DBTX, logger *zap.Logger) *postgresRepository {
	return &postgresRepository{
		queries: sqlc.New(qdb),
		logger:  logger,
	}
}

func (r *postgresRepository) GetPosts(ctx context.Context) ([]entity.Post, error) {
	posts, err := r.queries.GetPosts(ctx)
	if err != nil {
		return nil, convertError(err)
	}
	return convertRowsToEntity(posts), nil
}

func (r *postgresRepository) CreatePost(ctx context.Context, post entity.Post) (*entity.Post, error) {
	created, err := r.queries.CreatePost(ctx, sqlc.CreatePostParams{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Tags:     post.Tags,
	})
	if err != nil {
		return nil, convertError(err)
	}
	return convertRowToEntity(created), nil
}

func (r *postgresRepository) UpdatePostByID(ctx context.Context, id int, post entity.Post) (*entity.Post, error) {
	updated, err := r.queries.UpdatePostByID(ctx, sqlc.UpdatePostByIDParams{
		ID:       int32(id),
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Tags:     post.Tags,
	})
	if err != nil {
		return nil, convertError(err)
	}
	return convertRowToEntity(updated), nil
}

func (r *postgresRepository) GetPostByID(ctx context.Context, id int) (*entity.Post, error) {
	post, err := r.queries.GetPostByID(ctx, int32(id))
	if err != nil {
		return nil, convertError(err)
	}
	return convertRowToEntity(post), nil
}

func (r *postgresRepository) DeletePostByID(ctx context.Context, id int) (*entity.Post, error) {
	post, err := r.queries.DeletePostByID(ctx, int32(id))
	if err != nil {
		return nil, convertError(err)
	}
	return convertRowToEntity(post), nil
}

func convertRowsToEntity(rows []sqlc.Post) []entity.Post {
	p := make([]entity.Post, 0, len(rows))
	for _, v := range rows {
		p = append(p, *convertRowToEntity(v))
	}
	return p
}

func convertRowToEntity(v sqlc.Post) *entity.Post {
	return &entity.Post{
		ID:        int(v.ID),
		Title:     v.Title,
		Content:   v.Content,
		Category:  v.Category,
		Tags:      v.Tags,
		CreatedAt: v.Createdat.Time,
		UpdatedAt: v.Updatedat.Time,
	}
}

func convertError(err error) error {
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return errorspkg.ErrPostNotFound
	default:
		return err
	}
}
