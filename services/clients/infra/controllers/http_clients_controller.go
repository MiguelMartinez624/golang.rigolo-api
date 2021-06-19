package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/common"
	"github.com/rigolo-api/services/clients/domain"
	"log"
	"net/http"
)

type ClientsHTTPController struct {
	clientsDomain *domain.Domain
}

func NewClientsHTTPController(clientsDomain *domain.Domain) *ClientsHTTPController {
	return &ClientsHTTPController{clientsDomain: clientsDomain}
}

func (ctrl *ClientsHTTPController) CreateClient(c *gin.Context) {
	var newClient domain.NewClient
	if err := c.ShouldBindJSON(&newClient); err != nil {
		log.Printf("error binding json %+v", err)
		c.JSON(http.StatusBadRequest, common.Result{Error: err, Data: nil})
		return
	}
	_, err := ctrl.clientsDomain.RegisterNewClient(newClient)
	if err != nil {
		log.Printf("[RegisterNewClient] error :: %+v \n", err)
		c.JSON(http.StatusBadRequest, common.Result{Error: err, Data: nil})
		return
	}

	c.JSON(http.StatusOK, common.Result{Error: nil, Data: "client registed succefully"})
}
