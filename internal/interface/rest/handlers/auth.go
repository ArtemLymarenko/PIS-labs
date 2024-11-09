package handlers

import "github.com/gin-gonic/gin"

type AuthHandlerImpl interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Login godoc
//
//	@Summary		User login
//	@Description	User login with provided credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		dto.LoginRequest	true	"User login credentials"
//	@Success		200				{object}	dto.AuthResponse	"Successfully logged in"
//	@Failure		400				{object}	dto.ResponseErr		"Invalid credentials"
//	@Failure		500				{object}	dto.ResponseErr		"Internal Server Error"
//	@Router			/auth/login [post]
func (a *AuthHandler) Login(c *gin.Context) {}

// Register godoc
//
//	@Summary		User registration
//	@Description	User registration with required information
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		dto.RegisterRequest	true	"User registration details"
//	@Success		201				{object}	dto.AuthResponse	"User successfully registered"
//	@Failure		400				{object}	dto.ResponseErr		"Invalid data"
//	@Failure		500				{object}	dto.ResponseErr		"Internal Server Error"
//	@Router			/auth/register [post]
func (a *AuthHandler) Register(c *gin.Context) {}
