package integration_tests

import (
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/common/tests/containers"
	"github.com/rigolo-api/services/authentication"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	server            *gin.Engine = gin.Default()
	postgresContainer *containers.PostgresContainer
)

func TestMain(m *testing.M) {
	// Mount DB
	postgresContainer = containers.NewPostgresContainer()
	db, err := gorm.Open(postgres.Open(postgresContainer.URI), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}
	server = gin.Default()
	authentication.Bootstrap(db, server)

	// exit on test finish
	os.Exit(m.Run())
}
