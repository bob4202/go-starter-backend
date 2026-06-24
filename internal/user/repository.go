package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User

	query := `
	SELECT id, name, email, password_hash,is_active,created_at,updated_at,deleted_at
    FROM users WHERE email = $1 AND deleted_at IS NULL 
    `

	err := r.db.Get(&user, query, email)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
