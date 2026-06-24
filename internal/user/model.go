package user

import "time"

type User struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"hash_password" json:"-"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}
