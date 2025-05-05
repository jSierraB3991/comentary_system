package comentaryserviceinterface

import (
	"context"

	comentarymodels "github.com/jSierraB3991/comentary_system/domain/comentary_models"
)

type ComentarySystemServiceInterface interface {
	SaveComentary(ctx context.Context, comentary, commentUpId string, idPost, idUserComment uint) error
	GetComentarysByUser(ctx context.Context, userId uint) ([]comentarymodels.ComentarySystemComentaryModel, error)
}
