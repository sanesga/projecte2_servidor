package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/danilopolani/gocialite"
	"github.com/proyecto/backend_go/common"
	"gopkg.in/gin-gonic/gin.v1"
)

var gocial = gocialite.NewDispatcher()

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UsersSocialLogin(router *gin.RouterGroup) {
	fmt.Printf("entramos a routers.go")
	//router.GET("/", indexHandler)
	router.GET("/auth/:provider", redirectHandler)
	router.GET("/auth/:provider/callback", callbackHandler)

}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	router.POST("/:username/follow", ProfileFollow)
	router.DELETE("/:username/follow", ProfileUnfollow)
}

func ProfileRetrieve(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	profileSerializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": profileSerializer.Response()})
}

func ProfileFollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(UserModel)
	err = myUserModel.following(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func ProfileUnfollow(c *gin.Context) {
	username := c.Param("username")
	userModel, err := FindOneUser(&UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
		return
	}
	myUserModel := c.MustGet("my_user_model").(UserModel)

	err = myUserModel.unFollowing(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := ProfileSerializer{c, userModel}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

}

func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&UserModel{Email: loginValidator.userModel.Email})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}
	UpdateContextUserModel(c, userModel.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserRetrieve(c *gin.Context) {
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	myUserModel := c.MustGet("my_user_model").(UserModel)
	userModelValidator := NewUserModelValidatorFillWith(myUserModel)
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModelValidator.userModel.ID = myUserModel.ID
	if err := myUserModel.Update(userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	UpdateContextUserModel(c, myUserModel.ID)
	serializer := UserSerializer{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

/************************************social login********************************/

//Show homepage with login URL
// func indexHandler(c *gin.Context) {

// 	c.Writer.Write([]byte("<html><head><title>Gocialite example</title></head><body>" +
// 		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
// 		"</body></html>"))
// }

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
	fmt.Printf("estamos en redirectHandler")
	// Retrieve provider from route
	provider := c.Param("provider")

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     "b9563aec19bb264601a1",
			"clientSecret": "6c5cd9388386a6461a007576f4bfba1a7d144408",
			"redirectURL":  "http://localhost:8090/api/socialLogin/auth/github/callback",
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
	fmt.Printf("estamos en callback")
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

	userModel, err := FindOneUser(&UserModel{Username: user.Username})

	if err != nil {
		fmt.Printf("entra al if")

		// 	// c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		// 	// return
		// 	//SI EL USUARIO NO EXISTE, LO REGISTRAMOS

		// 	//ESTRUCTURA DE USERMODEL

		// type UserModel struct {
		// 	ID           uint    `gorm:"primary_key"`
		// 	Username     string  `gorm:"column:username"`
		// 	Email        string  `gorm:"column:email;unique_index"`
		// 	Bio          string  `gorm:"column:bio;size:1024"`
		// 	Image        *string `gorm:"column:image"`
		// 	PasswordHash string  `gorm:"column:password;not null"`
		// }

		// 	// userModel.ID = se asigna automáticamente
		userModel.Username = user.Username
		userModel.Email = user.Email
		userModel.Bio = user.FullName
		//userModel.Image = user.Avatar
		userModel.Image = nil
		userModel.PasswordHash = "12345678"

		socialLoginValidator := NewSocialLoginValidator()

		if err := socialLoginValidator.Bind(c); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
			return
		}
		fmt.Printf("tras el social login validator")

		fmt.Printf("%#v", userModel)

		if err := SaveOne(userModel); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}
		fmt.Printf("tras el save One")
		// 	// c.Set("my_user_model", socialLoginValidator.userModel)
		// 	// serializer := UserSerializer{c}
		// 	// c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
		// 	fmt.Printf("imprimo el contexto de gin")
		// 	fmt.Printf("%#", c)

		// 	if err := SaveOne(userModel); err != nil {
		// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		// 		return
		// 	}
		//guardem en contexte de gin
		// 	// c.Set("my_user_model", socialLoginValidator.userModel)
		// 	// serializer := UserSerializer{c}
		// 	// c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

		// 	//aquí hacemos el redirect

	} else {
		//guardamos el usuario en base de datos

		// if err := SaveOne(userModel); err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		// 	return
		// }
		// //guardamos usuario en contexto de gin
		// c.Set("my_user_model", userModel)
		// serializer := UserSerializer{c}
		// c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

		// socialLoginValidator := NewSocialLoginValidator()

		// if err := socialLoginValidator.Bind(c); err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		// 	return
		// }

		// if err := SaveOne(&socialLoginValidator.userModel); err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		// 	return
		// }
		// c.Set("my_user_model", socialLoginValidator.userModel)
		// serializer := UserSerializer{c}
		// c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})

	}

}

/**********************************************social login*************************************/
