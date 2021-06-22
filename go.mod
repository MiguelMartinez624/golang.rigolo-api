module github.com/rigolo-api

// +heroku goVersion go1.13
go 1.13

replace (
	github.com/rigolo-api/common v1.0.0 => ./common
	github.com/rigolo-api/services/authentication v1.0.0 => ./services/authentication
	github.com/rigolo-api/services/clients v1.0.0 => ./services/clients
)

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/google/uuid v1.2.0
	github.com/rigolo-api/common v1.0.0
	github.com/rigolo-api/services/authentication v1.0.0
	github.com/rigolo-api/services/clients v1.0.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)
