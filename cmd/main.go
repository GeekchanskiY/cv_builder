package main

import (
	"log"
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/config"
	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	database "github.com/GeekchanskiY/cv_builder/internal/db"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/router"
	server "github.com/GeekchanskiY/cv_builder/internal/server"

	"go.uber.org/fx"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(
			server.NewHTTPServer,
			router.CreateRoutes,
			database.GetDB,

			repository.CreateEmployeeRepository,
			repository.CreateDomainRepository,
			repository.CreateSkillRepository,
			repository.CreateCompanyRepository,

			controllers.CreateEmployeeController,
			controllers.CreateDomainController,
			controllers.CreateSkillController,
			controllers.CreateComapnyController,
		),
		fx.Invoke(
			func(*http.Server, *repository.EmployeeRepository) {},
			router.CreateEmployeeRoutes,
			router.CreateDomainRoutes,
			router.CreateSkillRoutes,
			router.CreateCompanyRoutes,
			config.SetupLogger,
		),
	)
}

func main() {

	// Connect to database
	connection, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	fx.New(CreateApp()).Run()

}
