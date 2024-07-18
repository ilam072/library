package users

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
)

const (
	ErrUniqueViolationCode = "23505"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(user model.User) error {
	const op = "repo.Users.Create"

	query, err := r.db.Prepare(`
	INSERT INTO users(username, password, email)
	VALUES ($1, $2, $3)
	RETURNING user_id
    `)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if _, err := query.Exec(user.Username, user.PasswordHash, user.Email); err != nil {
		if pqError, ok := err.(*pq.Error); ok && pqError.Code == ErrUniqueViolationCode {
			return myerrors.ErrUsernameExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) Get(inputUsername string) (model.User, error) {
	const op = "repo.Users.Get"

	query, err := r.db.Prepare(`
	SELECT user_id, username, password, email
	FROM users
	WHERE username = $1
	`)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	var (
		userId                        int
		username, passwordHash, email string
	)

	if err := query.QueryRow(inputUsername).Scan(&userId, &username, &passwordHash, &email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, myerrors.ErrInvalidUserData
		}
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return model.User{
		ID:           userId,
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}, nil
}

func (r *Repo) GetUserId(username string) (int, error) {
	const op = "repo.Users.GetUserId"

	query, err := r.db.Prepare(`
	SELECT user_id
	FROM users
	WHERE username = $1
	`)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var userId int

	if err := query.QueryRow(username).Scan(&userId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}

func (r *Repo) Update(username string, input model.User) error {
	const op = "repo.Users.Update"

	query, err := r.db.Prepare(`
	UPDATE users SET username = $1, password = $2, email = $3
	WHERE username = $4
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := query.Exec(input.Username, input.PasswordHash, input.Email, username); err != nil {
		if pqError, ok := err.(*pq.Error); ok && pqError.Code == ErrUniqueViolationCode {
			return myerrors.ErrUsernameExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) Delete(username string) error {
	const op = "repo.Users.Delete"

	query, err := r.db.Prepare(`
	DELETE FROM users
	WHERE username = $1
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := query.Exec(username); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
