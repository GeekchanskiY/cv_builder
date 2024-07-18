package main

import (
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal/config"
	"github.com/GeekchanskiY/cv_builder/internal/controllers"
	database "github.com/GeekchanskiY/cv_builder/internal/db"
	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/router"
	"github.com/GeekchanskiY/cv_builder/internal/server"

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
			repository.CreateCVRepository,
			repository.CreateProjectRepository,
			repository.CreateResponsibilityRepository,

			controllers.CreateEmployeeController,
			controllers.CreateDomainController,
			controllers.CreateSkillController,
			controllers.CreateCompanyController,
			controllers.CreateVacancyController,
			controllers.CreateCVController,
			controllers.CreateProjectController,
			controllers.CreateResponsibilityController,
			controllers.CreateUtilsController,
		),
		fx.Invoke(
			config.SetupLogger,
			router.CreateEmployeeRoutes,
			router.CreateDomainRoutes,
			router.CreateSkillRoutes,
			router.CreateCompanyRoutes,
			router.CreateVacancyRoutes,
			router.CreateCVRoutes,
			router.CreateProjectRoutes,
			router.CreateResponsibilityRoutes,
			func(*http.Server, *repository.EmployeeRepository) {},
		),
	)
}

func main() {
	config.LoadConfig()
	fx.New(CreateApp()).Run()
}
