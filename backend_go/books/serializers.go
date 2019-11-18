package books

import (
	"github.com/gosimple/slug"
	"gopkg.in/gin-gonic/gin.v1"
)

type BookSerializer struct {
	C *gin.Context
	BookModel
}

type BookResponse struct {
	ID          uint   `json:"-"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Author      string `json:"author"`
	Price       uint   `json:"price"`
}

type BooksSerializer struct {
	C     *gin.Context
	Books []BookModel
}

func (s *BookSerializer) Response() BookResponse {
	response := BookResponse{
		ID:          s.ID,
		Slug:        slug.Make(s.Title),
		Title:       s.Title,
		Description: s.Description,
		Category:    s.Category,
		Author:      s.Author,
		Price:       s.Price,
	}
	return response
}

func (s *BooksSerializer) Response() []BookResponse {
	response := []BookResponse{}
	for _, book := range s.Books {
		serializer := BookSerializer{s.C, book}
		response = append(response, serializer.Response())
	}
	return response
}
