package model

type Post struct {
	PostID      int     `json:"post_id" db:"post_id"`
	UserID      *int    `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	CreatedAt   *string `json:"created_at" db:"created_at"`
	UpdatedAt   *string `json:"updated_at" db:"updated_at"`
	// Comments    []*Comment `json:"comments" db:"comments"`
}
