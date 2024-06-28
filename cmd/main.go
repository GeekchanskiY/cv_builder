package main

import (
	"log"
	"net/http"

	"github.com/GeekchanskiY/cv_builder/internal"
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
			controllers.CreateEmployeeController,
		),
		fx.Invoke(
			func(*http.Server, *repository.EmployeeRepository) {},
			router.CreateEmployeeRoutes,
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
	internal.Samples(connection)

	log.Println("Server starting...")

	// router := router.CreateRoutes()
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("server_port")), router))

	fx.New(CreateApp()).Run()

}
