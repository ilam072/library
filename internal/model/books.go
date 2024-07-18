package model

type Book struct {
	BookID    int
	Title     string
	Author    string
	IssueYear int
	UserId    int
	FileName  string
	Access    string
}

type CreateBookDTO struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	IssueYear int    `json:"issue_year" binding:"required"`
	FileName  string `json:"file_name" binding:"required"`
	Access    string `json:"access" binding:"required"`
}

type GetBookDTO struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	IssueYear int    `json:"issue_year"`
}

type DownloadBookDTO struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	IssueYear int    `json:"issue_year"`
	FileName  string `json:"file_name"`
	Access    string `json:"access"`
}

type UpdateBookDTO struct {
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	IssueYear int    `json:"issue_year" binding:"required"`
	Access    string `json:"access" binding:"required"`
}
