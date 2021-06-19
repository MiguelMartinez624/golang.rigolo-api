package clients

import (
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/services/clients/domain"
	"github.com/rigolo-api/services/clients/infra/controllers"
	"github.com/rigolo-api/services/clients/infra/persistency"
	"gorm.io/gorm"
)

type ClientsServiceConfiguration struct {
	DB     *gorm.DB
	Server *gin.Engine
}

func Bootstrap(config ClientsServiceConfiguration) *domain.Domain {
	config.DB.Debug().AutoMigrate(domain.Client{})

	clientsRepo := persistency.NewPostgresClientsRepository(config.DB)
	clientsDomain := domain.NewDomain(clientsRepo)

	httpCtrl := controllers.NewClientsHTTPController(clientsDomain)

	//routing
	r := config.Server
	router := r.Group("clients")
	router.POST("", httpCtrl.CreateClient)
	router.GET("", httpCtrl.GetAllClients)

	return nil
}
