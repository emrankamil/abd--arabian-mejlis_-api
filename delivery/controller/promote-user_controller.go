package controller

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromoteController struct {
	promoteUsecase domain.PromoteUsecase
	userUsecase domain.UserUsecase
}

func NewPromoteController(uu domain.UserUsecase,pu domain.PromoteUsecase, env *bootstrap.Env) *PromoteController {
	return &PromoteController{
		promoteUsecase: pu,
		userUsecase: uu,
	}
}

func (pc *PromoteController) PromoteUser(c *gin.Context) {
	var userID = c.Param("id")
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
		return
	}

	admin, err := pc.userUsecase.GetUserByEmail(c, email.(string))
	if err != nil || admin.User_type != "ADMIN"{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
		return
	}

	err = pc.promoteUsecase.PromoteUser(c, userID)
	if err != nil{
		if err.Error() == "user with the given userID is not found"{
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
			return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true})
}


func (pc *PromoteController) DemoteUser(c *gin.Context) {
	var userID = c.Param("id")
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
		return
	}

	admin, err := pc.userUsecase.GetUserByEmail(c, email.(string))
	if err != nil || admin.User_type != "ADMIN"{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Unauthorized"})
		return
	}

	err = pc.promoteUsecase.DemoteUser(c, userID)
	if err != nil{
		if err.Error() == "user with the given userID is not found"{
			c.JSON(http.StatusNotFound, domain.ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
			return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true})
}