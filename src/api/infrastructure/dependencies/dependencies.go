package dependencies

import (
	"database/sql"

	"github.com/ExpertizeEafit/Api/src/api/infrastructure/configuration"

	_ "github.com/go-sql-driver/mysql"
)

// Container container for dependencies injection in modules
type Container struct {
	database *sql.DB
}

func StartDependencies() *Container {
	//config init
	config := configuration.Configuration()

	database, err := sql.Open(config.DBDriverName(), config.DatabaseConnectionString())
	if err != nil {
		panic(err.Error())
	}

	container := &Container{
		database: database,
	}

	return container
}

func (container *Container) Database() *sql.DB {
	return container.database
}
