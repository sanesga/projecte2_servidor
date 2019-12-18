package books

import (
	"github.com/gosimple/slug"
	"github.com/proyecto/backend_go/users"
	"github.com/gin-gonic/gin"
)

type BookSerializer struct {
	C *gin.Context
	BookModel
}

type BookResponse struct {
	ID             uint   `json:"-"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	Category       string `json:"category"`
	Author         string `json:"author"`
	Price          uint   `json:"price"`
	Favorite       bool   `json:"favorited"`
	FavoritesCount uint   `json:"favoritesCount"`
}

type BooksSerializer struct {
	C     *gin.Context
	Books []BookModel
}

func (s *BookSerializer) Response() BookResponse {
	myUserModel := s.C.MustGet("my_user_model").(users.UserModel)
	response := BookResponse{
		ID:             s.ID,
		Slug:           slug.Make(s.Title),
		Title:          s.Title,
		Description:    s.Description,
		Category:       s.Category,
		Author:         s.Author,
		Price:          s.Price,
		Favorite:       s.isFavoriteBy(GetBookUserModel(myUserModel)),
		FavoritesCount: s.favoritesCount(),
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

type BookUserSerializer struct {
	C *gin.Context
	BookUserModel
}

func (s *BookUserSerializer) Response() users.ProfileResponse {
	response := users.ProfileSerializer{s.C, s.BookUserModel.UserModel}
	return response.Response()
}

type CommentSerializer struct {
	C *gin.Context
	CommentBookModel
}
type CommentsSerializer struct {
	C        *gin.Context
	Comments []CommentBookModel
}
type CommentResponse struct {
	ID        uint                  `json:"id"`
	Body      string                `json:"body"`
	CreatedAt string                `json:"createdAt"`
	UpdatedAt string                `json:"updatedAt"`
	Author    users.ProfileResponse `json:"author"`
}

func (s *CommentSerializer) Response() CommentResponse {
	authorSerializer := BookUserSerializer{s.C, s.Author}
	response := CommentResponse{
		ID:        s.ID,
		Body:      s.Body,
		CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Author:    authorSerializer.Response(),
	}
	return response
}
func (s *CommentsSerializer) Response() []CommentResponse {
	response := []CommentResponse{}
	for _, comment := range s.Comments {
		serializer := CommentSerializer{s.C, comment}
		response = append(response, serializer.Response())
	}
	return response
}
