package dependencies

import (
	"details-microservice/configuration"
	"details-microservice/internal/infrastrcuture/repository/storage"
	"details-microservice/internal/infrastrcuture/rest/handler/reader"
	"details-microservice/internal/service/details"
)

type Dependencies struct {
	ReaderHandler reader.ReaderHandler
}

func InitDependencies(config *configuration.Configuration) Dependencies {
	// repository layer
	mapRepo := storage.NewMapRepository(config)

	// service layer
	detailsService := details.NewService(mapRepo)

	// handler layer
	readerHandler := reader.NewHandler(detailsService)

	return Dependencies{
		ReaderHandler: *readerHandler,
	}
}
