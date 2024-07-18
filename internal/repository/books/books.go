package books

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"islamic-library/internal/model"
	"islamic-library/internal/myerrors"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(book model.Book) error {
	const op = "repo.Books.Create"

	query, err := r.db.Prepare(`
	INSERT INTO books(title, author, issue_year, user_id, file_name, book_access)
	VALUES ($1, $2, $3, $4, $5, $6)
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := query.Exec(
		book.Title, book.Author, book.IssueYear, book.UserId, book.FileName, book.Access,
	); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAll(username string) ([]model.Book, error) {
	const op = "repo.Books.GetAll"

	query, err := r.db.Prepare(`
	SELECT book_id, title, author, issue_year, user_id, file_name, book_access
	FROM books
	INNER JOIN users USING(user_id)
	WHERE username = $1
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var books []model.Book

	rows, err := query.Query(username)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.IssueYear,
			&book.UserId,
			&book.FileName,
			&book.Access,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, myerrors.ErrBooksNotFound
	}

	return books, nil
}

func (r *Repo) GetPublic(username string) ([]model.Book, error) {
	const op = "repo.Books.GetPublic"

	query, err := r.db.Prepare(`
	SELECT book_id, title, author, issue_year, user_id, file_name, book_access
	FROM books
	INNER JOIN users USING(user_id)
	WHERE username = $1 AND book_access = 'public'
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var books []model.Book

	rows, err := query.Query(username)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.IssueYear,
			&book.UserId,
			&book.FileName,
			&book.Access,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, myerrors.ErrBooksNotFound
	}

	return books, nil
}

func (r *Repo) GetById(bookId int) (model.Book, error) {
	const op = "repo.Books.GetById"

	query, err := r.db.Prepare(`
	SELECT book_id, title, author, issue_year, user_id, file_name, book_access
	FROM books
	WHERE book_id = $1
	`)
	if err != nil {
		return model.Book{}, fmt.Errorf("%s: %w", op, err)
	}

	var book model.Book

	if err := query.QueryRow(bookId).Scan(
		&book.BookID,
		&book.Title,
		&book.Author,
		&book.IssueYear,
		&book.UserId,
		&book.FileName,
		&book.Access,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, myerrors.ErrBookNotFound
		}
		return model.Book{}, fmt.Errorf("%s: %w", op, err)
	}

	return book, nil
}

func (r *Repo) Update(bookId int, username string, input model.Book) error {
	const op = "repo.Books.Update"

	query, err := r.db.Prepare(`
	UPDATE books AS b
    SET title = $1,
	    author = $2,
	    issue_year = $3,
	    book_access = $4
    FROM users AS u	
    WHERE b.user_id = u.user_id and u.username = $5 and b.book_id = $6
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := query.Exec(
		input.Title,
		input.Author,
		input.IssueYear,
		input.Access,
		username,
		bookId,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return myerrors.ErrBookNotFound
	}

	return nil
}

func (r *Repo) Delete(bookId int, username string) error {
	const op = "repo.Books.Delete"

	query, err := r.db.Prepare(`
	DELETE FROM books
	USING users
	WHERE books.user_id = users.user_id AND book_id = $1 AND username = $2
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := query.Exec(bookId, username); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
