package main

import (
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
			repository.CreateVacanciesRepository,

			controllers.CreateEmployeeController,
			controllers.CreateDomainController,
			controllers.CreateSkillController,
			controllers.CreateComapnyController,
			controllers.CreateVacancyController,
		),
		fx.Invoke(
			func(*http.Server, *repository.EmployeeRepository) {},
			router.CreateEmployeeRoutes,
			router.CreateDomainRoutes,
			router.CreateSkillRoutes,
			router.CreateCompanyRoutes,
			router.CreateVacancyRoutes,
			config.SetupLogger,
		),
	)
}

func main() {
	fx.New(CreateApp()).Run()
}
