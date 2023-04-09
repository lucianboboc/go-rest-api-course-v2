//go:build integration

package db

import (
	"context"
	"testing"

	"github.com/lucianboboc/go-rest-api-course/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		DB, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := DB.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := DB.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		DB, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := DB.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		err = DB.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = DB.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
