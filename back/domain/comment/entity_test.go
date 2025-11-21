package comment_test

import (
	"strings"
	"testing"

	"poc2/back/domain/comment"

	"github.com/stretchr/testify/assert"
)

func TestNewCommentContent(t *testing.T) {
	t.Run("valid content", func(t *testing.T) {
		content, err := comment.NewCommentContent("This is a valid comment")
		assert.NoError(t, err)
		assert.Equal(t, "This is a valid comment", content.String())
	})

	t.Run("empty content returns error", func(t *testing.T) {
		_, err := comment.NewCommentContent("")
		assert.Error(t, err)
		assert.Equal(t, "comment content cannot be empty", err.Error())
	})

	t.Run("too long content returns error", func(t *testing.T) {
		longString := strings.Repeat("a", 501)
		_, err := comment.NewCommentContent(longString)
		assert.Error(t, err)
		assert.Equal(t, "comment content is too long", err.Error())
	})
}
