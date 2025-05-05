package comentaryrepository

import (
	"context"

	comentarymodels "github.com/jSierraB3991/comentary_system/domain/comentary_models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) SaveComentary(ctx context.Context, comentary comentarymodels.ComentarySystemComentaryModel, comentaryUpId string) error {
	if comentaryUpId != "" {
		filter := bson.M{"id": comentaryUpId}
		update := bson.M{
			"$push": bson.M{"coments": comentary},
		}

		_, err := r.GetComentaryCollection().UpdateOne(ctx, filter, update)
		return err
	} else {
		_, err := r.GetComentaryCollection().InsertOne(ctx, comentary)
		return err
	}
}

func (r *Repository) GetComentarysByUser(ctx context.Context, userId uint) ([]comentarymodels.ComentarySystemComentaryModel, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"id_user": userId},         // Comentarios principales del usuario
			{"coments.id_user": userId}, // Comentarios hijos del usuario
		},
	}

	cursor, err := r.GetComentaryCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultados []comentarymodels.ComentarySystemComentaryModel
	if err := cursor.All(ctx, &resultados); err != nil {
		return nil, err
	}
	return resultados, nil
}
