package develop

import (
	"fmt"
	"os"

	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application"
)

type Configuration struct{}

const (
	dbLocalUsername = "local_username"
	dbLocalPassword = "local_password"
	dbLocalHost     = "localhost"
	dbName          = "expertize"
)

func (*Configuration) DatabaseConnectionString() string {
	var username string
	var password string
	var host string
	if application.IsLocal() {
		username = os.Getenv(dbLocalUsername)
		password = os.Getenv(dbLocalPassword)
		host = dbLocalHost
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, dbName)
	return connection
}
