package production

import (
	"fmt"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application"
	"os"
)

type Configuration struct{}

const (
	dbUsername = ""
	dbPassword = ""
	dbHost     = ""
	dbName     = "expertize"
)

func (*Configuration) DatabaseConnectionString() string {
	var username string
	var password string
	var host string
	if application.IsLocal() {
		username = dbUsername
		password = os.Getenv(dbPassword)
		host = dbHost
	}
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", username, password, host, dbName)
	return connection
}
