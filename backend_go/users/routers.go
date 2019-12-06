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
func VerTodos(router *gin.RouterGroup) {
	fmt.Printf("entra a ver todos")
	router.GET("/", VerUsers) /*llama al método BookList (más abajo*/
}

//Rutas para hacer el social login
func UsersSocialLogin(router *gin.RouterGroup) {
	router.GET("/auth/:provider", redirectHandler)
	router.GET("/auth/:provider/callback", callbackHandler)
}

//Ruta parar recuperar los datos del usuario, tras hacer el social login
func UserSocial(router *gin.RouterGroup) {
	router.GET("/:username", getUser)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
	//router.DELETE("/", UserDelete) //NO FUNCIONA TODAVÍA
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)
	router.POST("/:username/follow", ProfileFollow)
	router.DELETE("/:username/follow", ProfileUnfollow)
}
func VerUsers(c *gin.Context) {
	userModels, err := getAll()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("users", errors.New("Invalid param")))
		return
	}
	serializer := UsersSerializer{c, userModels}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
func getUser(c *gin.Context) {
	//OBTENER EL MAIL DEL USUARIO DE LA BASE DE DATOS
	//le pasamos el slug por la petición get desde el frontend
	slug := c.Param("slug")
	fmt.Printf("%#v", slug)
	//buscamos el usuario
	userModel, err := FindOneUser(&UserModel{Username: slug})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("user", errors.New("Invalid username")))
		return
	}
	//guardamos en contexto de gin
	c.Set("my_user_model", userModel)
	//le pasamos nuestro propio serializer que devuelve solo el email (lo necesitamos para hacer login)
	serializer := socialUserSerializer{c, userModel}
	//el serializer nos devuelve los datos
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
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
	fmt.Printf("estamos en users login")
	loginValidator := NewLoginValidator()

	if err := loginValidator.Bind(c); err != nil {
		fmt.Printf("estamos dentro del validator")
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	fmt.Printf("despues del validator")

	userModel, err := FindOneUser(&UserModel{Email: loginValidator.userModel.Email})

	fmt.Printf("%#v", userModel)

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

	fmt.Printf("fin del login")
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

// func UserDelete(c *gin.Context) {
// 	myUserModel := c.MustGet("my_user_model").(UserModel)
// 	userModelValidator := NewUserModelValidatorFillWith(myUserModel)
// 	if err := userModelValidator.Bind(c); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
// 		return
// 	}

// 	userModelValidator.userModel.ID = myUserModel.ID
// 	if err := myUserModel.Delete(userModelValidator.userModel); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
// 		return
// 	}
// 	UpdateContextUserModel(c, myUserModel.ID)
// 	serializer := UserSerializer{c}
// 	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
// }

/************************************social login********************************/

// Redirect to correct oAuth URL
func redirectHandler(c *gin.Context) {
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
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	//provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	//fmt.Printf("informacion del usuario")
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)
	//fmt.Printf("fin de la informacion de usuario")

	// If no errors, show user
	// c.Writer.Write([]byte("FullName: " + user.FullName + "\n"))
	// c.Writer.Write([]byte("Email: " + user.Email + "\n"))
	// c.Writer.Write([]byte("Username: " + user.Username + "\n"))
	// c.Writer.Write([]byte("Avatar: " + user.Avatar + "\n"))

	userModel, err := FindOneUser(&UserModel{Username: user.Username})

	if err != nil {
		//no se ha encontrado el usuario, no existe, hay que registrarlo

		//rellenamos los campos con los datos que nos llegan de github
		userModel.Username = user.Username
		userModel.Email = user.Email
		userModel.Bio = user.FullName
		userModel.Image = nil
		userModel.PasswordHash = "12345678"

		//lo guardamos en la base de datos
		if err := SaveOne(&userModel); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
			return
		}
		//guardamos en el contexto de gin
		c.Set("my_user_model", userModel)
		//hacemos el redirect
		c.Redirect(http.StatusFound, "http://localhost:8081/social"+userModel.Email)

	} else {

		//EL USUARIO EXISTE, HACEMOS LOGIN
		//guardamos en el contexto de gin
		c.Set("my_user_model", userModel)
		//hacemos el redirect
		c.Redirect(http.StatusFound, "http://localhost:8081/social/"+userModel.Email)
	}
}

/**********************************************social login*************************************/
