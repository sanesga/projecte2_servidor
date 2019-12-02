package main

import (
	"github.com/danilopolani/gocialite"
	"github.com/jinzhu/gorm"
	"github.com/proyecto/backend_go/articles"
	"github.com/proyecto/backend_go/books"
	"github.com/proyecto/backend_go/common"
	"github.com/proyecto/backend_go/users"

	"gopkg.in/gin-gonic/gin.v1"
)

/*social login*/
var gocial = gocialite.NewDispatcher()

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
	db.AutoMigrate(&books.BookModel{})
}

func main() {

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	MakeRoutes(r) //habilita els CORS

	/*no requiere token*/
	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	users.UsersSocialLogin(v1.Group("/socialLogin"))

	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	books.BooksAnonymousRegister(v1.Group("/books")) /*Esto llama a routes.go*/
	books.BooksRegister(v1.Group("/books"))

	/*sí requiere token*/
	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))
	articles.ArticlesRegister(v1.Group("/articles"))

	r.Run(":8090") // listen and serve on localhost:8090
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()

		/*
			fmt.Printf("c.Request.Method \n")
			fmt.Printf(c.Request.Method)
			fmt.Printf("c.Request.RequestURI \n")
			fmt.Printf(c.Request.RequestURI)
		*/
	}
	r.Use(cors)
}
