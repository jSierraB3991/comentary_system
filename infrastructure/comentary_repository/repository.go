package comentaryrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Client
}

func InitiateRepo(collection *mongo.Client) *Repository {
	return &Repository{
		collection: collection,
	}
}

func (r *Repository) GetComentaryCollection() *mongo.Collection {
	return r.collection.Database("comentary").Collection("comentary")
}
