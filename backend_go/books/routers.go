package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/proyecto/backend_go/common"
	"gopkg.in/gin-gonic/gin.v1"
)

func BooksRegister(router *gin.RouterGroup) {
	router.POST("/", BookCreate)
	router.PUT("/:slug", BookUpdate)
	router.DELETE("/:slug", BookDelete)
	router.POST("/:slug/comments", BookCommentCreate)
	router.DELETE("/:slug/comments/:id", BookCommentDelete)
}

func BooksAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", BookList) /*llama al método BookList (más abajo*/
	router.GET("/:slug", BookRetrieve)
	router.GET("/:slug/comments", BookCommentList)
}
func BookCommentCreate(c *gin.Context) {
	slug := c.Param("slug")
	bookModel, err := FindOneBook(&BookModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid slug")))
		return
	}
	commentModelValidator := NewCommentModelValidator()
	if err := commentModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	commentModelValidator.commentModel.Book = bookModel

	if err := SaveOne(&commentModelValidator.commentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := CommentSerializer{c, commentModelValidator.commentModel}
	c.JSON(http.StatusCreated, gin.H{"comment": serializer.Response()})
}
func BookCommentDelete(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid id")))
		return
	}
	err = DeleteCommentModel([]uint{id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("Invalid id")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment": "Delete success"})
}
func BookCommentList(c *gin.Context) {
	slug := c.Param("slug")
	bookModel, err := FindOneBook(&BookModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Invalid slug")))
		return
	}
	err = bookModel.getComments()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("Database error")))
		return
	}
	serializer := CommentsSerializer{c, bookModel.Comments}
	c.JSON(http.StatusOK, gin.H{"comments": serializer.Response()})
}

func BookCreate(c *gin.Context) {
	bookModelValidator := NewBookModelValidator()
	if err := bookModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	//fmt.Println(bookModelValidator.bookModel)

	if err := SaveOne(&bookModelValidator.bookModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := BookSerializer{c, bookModelValidator.bookModel}
	c.JSON(http.StatusCreated, gin.H{"book": serializer.Response()})
}

func BookList(c *gin.Context) {
	bookModels, err := getAllBooks()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("Invalid param")))
		return
	}
	serializer := BooksSerializer{c, bookModels}
	c.JSON(http.StatusOK, gin.H{"book": serializer.Response()})
}

func BookRetrieve(c *gin.Context) {
	slug := c.Param("slug")

	bookModel, err := FindOneBook(&BookModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("Invalid slug")))
		return
	}
	serializer := BookSerializer{c, bookModel}
	c.JSON(http.StatusOK, gin.H{"book": serializer.Response()})
}

func BookUpdate(c *gin.Context) {
	slug := c.Param("slug")
	bookModel, err := FindOneBook(&BookModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("Invalid slug")))
		return
	}
	bookModelValidator := NewBookModelValidatorFillWith(bookModel)
	if err := bookModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	bookModelValidator.bookModel.ID = bookModel.ID
	if err := bookModel.Update(bookModelValidator.bookModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := BookSerializer{c, bookModel}
	c.JSON(http.StatusOK, gin.H{"book": serializer.Response()})
}

func BookDelete(c *gin.Context) {
	slug := c.Param("slug")
	err := DeleteBookModel(&BookModel{Slug: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("books", errors.New("Invalid slug")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"book": "Delete success"})
}
