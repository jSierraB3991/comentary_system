package comentaryservice

import (
	"context"

	comentarylibs "github.com/jSierraB3991/comentary_system/domain/comentary_libs"
	comentaryrepositoryinterface "github.com/jSierraB3991/comentary_system/domain/comentary_repository_interface"
	comentaryserviceinterface "github.com/jSierraB3991/comentary_system/domain/comentary_service_interface"
	comentaryrepository "github.com/jSierraB3991/comentary_system/infrastructure/comentary_repository"
)

type ComentarySystemService struct {
	perspectiveService comentaryserviceinterface.PerspectiveApiServiceInterface
	repository         comentaryrepositoryinterface.MongoDbRepositoryInterface
}

func NewComentarySystemService() *ComentarySystemService {
	env := comentarylibs.NewEnviroment()
	perspectiveService := NewPerspectiveService(env.PerspectiveAPIKey, env.UrlBasePerspective)

	mongoClient := comentaryrepository.ConnectToMongo(env.MongoDbUri, context.Background())
	repository := comentaryrepository.InitiateRepo(mongoClient)

	return &ComentarySystemService{
		perspectiveService: perspectiveService,
		repository:         repository,
	}
}
