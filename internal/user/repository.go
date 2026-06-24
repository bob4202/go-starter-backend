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

func (r *Repository) FindById(id string) (*User, error) {
	var user User
	query := `
	SELECT id, name, email, password_hash,is_active,created_at,updated_at,deleted_at
    FROM users WHERE id  = $1 AND deleted_at IS NULL 
	`
	err := r.db.Get(&user, query, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateUser(u *User) error {

	query := `
	INSERT INTO users(name, email, password_hash) 
	VALUES ($1, $2, $3)
	RETURNING
		id, created_at, updated_at, is_active
	`

	return r.db.QueryRow(query, u.Name, u.Email, u.PasswordHash).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.IsActive)

}

func (r *Repository) UpdateUser(u *User) error {
	query := `
	UPDATE users 
	SET name=$2,email=$3,updated_at=NOW()
	WHERE
		id = $1
	RETURNING
		name,email, is_active,updated_at,deleted_at
	`

	return r.db.QueryRow(query, u.ID, u.Name, u.Email).Scan(&u.Name, &u.Email, &u.IsActive, &u.UpdatedAt)

}

func (r *Repository) ChangePassword(id, passwordHash string) error {
	query := `
	UPDATE users
	SET password=$2, updated_at=NOW()
	WHERE
		id = $1
	`

	_, err := r.db.Exec(query, id, passwordHash)

	if err != nil {
		return err
	}

	return nil

}

func (r *Repository) DeleteUser(id string) error {
	query := `
	UPDATE users
	SET deleted_at = NOW(), updated_at = NOW()
	WHERE
		id = $1
	AND deleted_at IS NULL
	RETURNING
		deleted_at
	`

	_, err := r.db.Exec(query, id)

	return err
}
