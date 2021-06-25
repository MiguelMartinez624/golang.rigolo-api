package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/services/authentication/domain"
	"github.com/rigolo-api/services/authentication/domain/entities"
	"github.com/rigolo-api/services/authentication/infra/controllers"
	"github.com/rigolo-api/services/authentication/infra/persistency"

	"gorm.io/gorm"
)

func Bootstrap(db *gorm.DB, r *gin.Engine) *domain.Context {
	//migrations
	db.Debug().AutoMigrate(entities.Account{})

	authRepo := persistency.BuildAccountRepositoryPostgres(db)
	boundedContext := domain.BuildContext(authRepo)
	ctrl := controllers.BuildAuthenticationHTTPCtrl(boundedContext)

	router := r.Group("auth")

	router.POST("/signup", ctrl.CreateAccount)
	router.POST("/signin", ctrl.AuthenticateAccount)

	return boundedContext
}
