package comentarylibs

import (
	"log"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

type Enviroment struct {
	PerspectiveAPIKey  string
	UrlBasePerspective string

	MongoDbUri string
}

func NewEnviroment() *Enviroment {
	perspectiveAPIKey, err := jsierralibs.GetDataOfEnviromentRequired("API_PERSPECTTIVE_KEY")
	if err != nil {
		log.Fatal(err)
	}

	perspectiveUrl, err := jsierralibs.GetDataOfEnviromentRequired("URL_PERSPECTIVE")
	if err != nil {
		log.Fatal(err)
	}
	mongoDbUri, err := jsierralibs.GetDataOfEnviromentRequired("MONGO_DB_URI")
	if err != nil {
		log.Fatal(err)
	}

	return &Enviroment{
		PerspectiveAPIKey:  perspectiveAPIKey,
		UrlBasePerspective: perspectiveUrl,
		MongoDbUri:         mongoDbUri,
	}
}
