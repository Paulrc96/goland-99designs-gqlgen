package model

type Comment struct {
	CommentID   int     `json:"comment_id" db:"comment_id"`
	PostID      *int    `json:"post_id" db:"post_id"`
	UserID      *int    `json:"user_id" db:"user_id"`
	Description *string `json:"description" db:"description"`
	CreatedAt   *string `json:"created_at" db:"created_at"`
	UpdatedAt   *string `json:"updated_at" db:"updated_at"`
}
