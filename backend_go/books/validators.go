package books

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/proyecto/backend_go/common"
	"github.com/proyecto/backend_go/users"
)

type BookModelValidator struct {
	Book struct {
		Title       string `form:"title" json:"title" binding:"min=4"`
		Description string `form:"description" json:"description" binding:"max=2048"`
		Category    string `form:"category" json:"category" binding:"max=2048"`
		Author      string `form:"author" json:"author" binding:"max=2048"`
		//no he puesto price

	} `json:"book"`
	bookModel BookModel `json:"-"`
}

func NewBookModelValidator() BookModelValidator {
	return BookModelValidator{}
}

func NewBookModelValidatorFillWith(bookModel BookModel) BookModelValidator {
	bookModelValidator := NewBookModelValidator()
	bookModelValidator.Book.Title = bookModel.Title
	bookModelValidator.Book.Description = bookModel.Description
	bookModelValidator.Book.Category = bookModel.Category
	bookModelValidator.Book.Author = bookModel.Author

	return bookModelValidator
}

func (s *BookModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.bookModel.Slug = slug.Make(s.Book.Title)
	s.bookModel.Title = s.Book.Title
	s.bookModel.Description = s.Book.Description
	s.bookModel.Category = s.Book.Category
	s.bookModel.Author = s.Book.Author
	return nil
}

type CommentBookModelValidator struct {
	Comment struct {
		Body string `form:"body" json:"body" binding:"max=2048"`
	} `json:"comment"`
	commentBookModel CommentBookModel `json:"-"`
}

func NewCommentBookModelValidator() CommentBookModelValidator {
	return CommentBookModelValidator{}
}

func (s *CommentBookModelValidator) Bind(c *gin.Context) error {
	myUserModel := c.MustGet("my_user_model").(users.UserModel)

	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.commentBookModel.Body = s.Comment.Body
	s.commentBookModel.Author = GetBookUserModel(myUserModel)
	return nil
}
