package main

import (
	"fmt"
	"net/http"

	"github.com/danilopolani/gocialite"
	"github.com/gin-gonic/gin"
)

var gocial = gocialite.NewDispatcher()

func main() {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/auth/:provider", redirectHandler)
	router.GET("/auth/:provider/callback", callbackHandler)

	router.Run("127.0.0.1:9091")
}

// Show homepage with login URL
func indexHandler(c *gin.Context) {
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
			"redirectURL":  "http://localhost:3003/auth/github/callback",
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
