package books

import (
	"errors"
	"net/http"

	"github.com/backend_go/common"
	"gopkg.in/gin-gonic/gin.v1"
)

func BooksRegister(router *gin.RouterGroup) {
	router.POST("/", BookCreate)
	router.PUT("/:slug", BookUpdate)
	router.DELETE("/:slug", BookDelete)
}

func BooksAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", BookList) /*llama al método BookList (más abajo*/
	router.GET("/:slug", BookRetrieve)
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
