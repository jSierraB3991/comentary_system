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
