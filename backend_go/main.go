package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/proyecto/backend_go/articles"
	"github.com/proyecto/backend_go/books"
	"github.com/proyecto/backend_go/common"
	"github.com/proyecto/backend_go/users"

	"gopkg.in/gin-gonic/gin.v1"

	/*for social login*/
	"net/http"

	"github.com/danilopolani/gocialite"
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

	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	books.BooksAnonymousRegister(v1.Group("/books")) /*Esto llama a routes.go*/
	books.BooksRegister(v1.Group("/books"))

	/*social login*/
	v1.Group("/socialLogin", indexHandler)
	v1.Group("/socialLogin/auth/:provider", redirectHandler)
	v1.Group("/socialLogin/auth/:provider/callback", callbackHandler)

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

/**********************social login********************************/

// Show homepage with login URL
func indexHandler(c *gin.Context) {
	fmt.Printf("entra a social login")
	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"</body></html>"))
}

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     "54b66a2c7590fd7df695",
			"clientSecret": "d4e933ea85b321b9c235372eb2a4a6c75293501d",
			"redirectURL":  "http://localhost:9091/auth/github/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{"public_repo"},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

// Handle callback of provider
func callbackHandler(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	// provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

	// If no errors, show user
	c.Writer.Write([]byte("FullName: " + user.FullName + "\n"))
	c.Writer.Write([]byte("Email: " + user.Email + "\n"))
	c.Writer.Write([]byte("Username: " + user.Username + "\n"))
	c.Writer.Write([]byte("Avatar: " + user.Avatar + "\n"))
}
