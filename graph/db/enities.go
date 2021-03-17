package database

type Comment struct {
	CommentID   *int    `json:"comment_id"`
	Description *string `json:"description"`
}

type Post struct {
	PostID      *int       `json:"post_id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Comments    []*Comment `json:"comments"`
}

type User struct {
	ID              int     `json:"id"`
	Name            *string `json:"name"`
	Email           *string `json:"email"`
	LastName        *string `json:"last_name" db:"last_name"`
	Birthday        *string `json:"birthday" db:"birthday"`
	Address         *string `json:"address" db:"address"`
	EmailVerifiedAt *string `json:"email_verified_at" db:"email_verified_at"`
	Password        *string `json:"password" db:"password"`
	RememberToken   *string `json:"remember_token" db:"remember_token"`
	CreatedAt       *string `json:"created_at" db:"created_at"`
	UpdatedAt       *string `json:"updated_at" db:"updated_at"`
	Posts           []*Post `json:"posts"`
}
