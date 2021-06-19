package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/services/authentication/domain"

	"log"
	"net/http"
)

type AuthenticationHTTPCtrl struct {
	authContext *domain.Context
}

func BuildAuthenticationHTTPCtrl(authContext *domain.Context) *AuthenticationHTTPCtrl {
	return &AuthenticationHTTPCtrl{authContext}
}

func (ctrl *AuthenticationHTTPCtrl) CreateAccount(c *gin.Context) {
	var newAccount data.NewAccount
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		log.Printf("error binding json %+v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := ctrl.authContext.RegisterAccount(newAccount)
	if err != nil {
		log.Printf("[RegisterAccount] error :: %+v \n", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctrl *AuthenticationHTTPCtrl) AuthenticateAccount(c *gin.Context) {
	var credentials data.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Printf("error binding json %+v", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := ctrl.authContext.AuthenticateAccount(credentials)
	if err != nil {
		log.Printf("[AuthenticateAccount] error :: %+v \n", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
