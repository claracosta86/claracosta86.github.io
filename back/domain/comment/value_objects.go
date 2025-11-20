package comment

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type CommentContent struct {
	value string
}

func NewCommentContent(value string) (CommentContent, error) {
	if value == "" {
		return CommentContent{}, errors.New("comment content cannot be empty")
	}
	if len(value) > 500 {
		return CommentContent{}, errors.New("comment content is too long")
	}
	return CommentContent{value: value}, nil
}

func (c CommentContent) String() string {
	return c.value
}

func (c CommentContent) Value() (driver.Value, error) {
	return c.value, nil
}

// Scan implements the sql.Scanner interface
func (c *CommentContent) Scan(value interface{}) error {
	if value == nil {
		c.value = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		c.value = string(v)
	case string:
		c.value = v
	default:
		return fmt.Errorf("failed to scan CommentContent: %v", value)
	}
	return nil
}
