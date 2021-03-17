package model

type Client struct {
	ID        int     `json:"id" db:"id"`
	Name      *string `json:"name" db:"name"`
	Email     *string `json:"email" db:"email"`
	LastName  *string `json:"last_name" db:"last_name"`
	Birthday  *string `json:"birthday" db:"birthday"`
	Address   *string `json:"address" db:"address"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	UpdatedAt *string `json:"updated_at" db:"updated_at"`
}
