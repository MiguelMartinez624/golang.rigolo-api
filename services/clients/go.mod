module github.com/rigolo-api/services/clients

go 1.13

replace github.com/rigolo-api/common => ../../common

require (
	github.com/Microsoft/hcsshim v0.8.16
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/rigolo-api/common v1.0.0
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/text v0.3.6 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)
