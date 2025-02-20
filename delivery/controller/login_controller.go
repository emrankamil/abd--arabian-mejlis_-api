package controller

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)


type LoginController struct {
	UserUsecase domain.UserUsecase
	LoginUsecase domain.LoginUsecase
	Env 		*bootstrap.Env
}

func NewLoginController(uu domain.UserUsecase, lu domain.LoginUsecase, env *bootstrap.Env) *LoginController {
	return &LoginController{
		UserUsecase: uu,
		LoginUsecase: lu,
		Env: env,
	}
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest,domain.ErrorResponse{Error:err.Error()})
	}

	user, err := lc.UserUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound,domain.ErrorResponse{Error:err.Error()})
	}

	
	if err := infrastructure.VerifyPassword(request.Password, user.Password); err != nil{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	updatedUser := &domain.User{
		ID: user.ID,
		Token: accessToken,
		Refresh_token: refreshToken,
	}

	err = lc.UserUsecase.UpdateUser(c, updatedUser)

	if err != nil{
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Data: loginResponse})
}

func (lc *LoginController) Logout(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
		return
	}

	err := lc.LoginUsecase.LogoutUser(c, email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true})
}