package controller

import (
	"easyapi/blog/internal/entity"
	errorspkg "easyapi/blog/internal/errors"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controller) getPosts(ctx *gin.Context) {
	posts, err := c.repo.GetPosts(ctx.Request.Context())

	if err != nil {
		convertError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func (c *controller) createPost(ctx *gin.Context) {
	var newPost entity.Post

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts, err := c.repo.CreatePost(ctx.Request.Context(), newPost)

	if err != nil {
		convertError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func (c *controller) updatePostByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var newPost entity.Post

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	iid, err := strconv.Atoi(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id should be a number"})
		return
	}

	posts, err := c.repo.UpdatePostByID(ctx.Request.Context(), iid, newPost)

	if err != nil {
		convertError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func (c *controller) getPostByID(ctx *gin.Context) {
	id := ctx.Param("id")

	iid, err := strconv.Atoi(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id should be a number"})
		return
	}

	post, err := c.repo.GetPostByID(ctx.Request.Context(), iid)

	if err != nil {
		convertError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, post)
}

func (c *controller) deletePostByID(ctx *gin.Context) {
	id := ctx.Param("id")

	iid, err := strconv.Atoi(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id should be a number"})
		return
	}

	posts, err := c.repo.DeletePostByID(ctx.Request.Context(), iid)

	if err != nil {
		convertError(ctx, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func convertError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, errorspkg.ErrPostNotFound):
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	default:
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
