package users

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/proyecto/backend_go/common"
)

//nos creamos nuestro propio serializer, pasándole el contexto y el usuario
type socialUserSerializer struct {
	C *gin.Context
	UserModel
}

//nos devuelve el mail
type socialUserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (self *socialUserSerializer) Response() socialUserResponse {
	//recupera el usuario del contexto de gin
	myUserModel := self.C.MustGet("my_user_model").(UserModel)
	user := socialUserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

// Put your response logic including wrap the userModel here.
func (self *ProfileSerializer) Response() ProfileResponse {
	myUserModel := self.C.MustGet("my_user_model").(UserModel)
	profile := ProfileResponse{
		ID:        self.ID,
		Username:  self.Username,
		Bio:       self.Bio,
		Image:     self.Image,
		Following: myUserModel.isFollowing(self.UserModel),
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}

type UsersSerializer struct {
	C     *gin.Context
	Users []UserModel
}

func (s *UsersSerializer) Response() []UsuarioResponse {
	response := []UsuarioResponse{}
	for _, user := range s.Users {
		serializer := UsuarioSerializer{s.C, user}
		fmt.Printf("c.Request.Method \n")
		response = append(response, serializer.Response())
	}
	return response
}

type UsuarioSerializer struct {
	C *gin.Context
	UserModel
}

type UsuarioResponse struct {
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Bio          string  `json:"bio"`
	Image        *string `json:"image"`
	Token        string  `json:"token"`
	PasswordHash string  `json:"password"`
}

func (s *UsuarioSerializer) Response() UsuarioResponse {
	response := UsuarioResponse{
		Username:     s.Username,
		Email:        s.Email,
		Bio:          s.Bio,
		Image:        s.Image,
		Token:        common.GenToken(s.ID),
		PasswordHash: s.PasswordHash,
	}
	return response
}
