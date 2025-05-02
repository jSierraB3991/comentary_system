package comentaryservice

import (
	"encoding/json"
	"fmt"

	comentarylibs "github.com/jSierraB3991/comentary_system/domain/comentary_libs"
	comentaryrequest "github.com/jSierraB3991/comentary_system/infrastructure/comentary_request"
	comentaryresponse "github.com/jSierraB3991/comentary_system/infrastructure/comentary_response"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

type PerspectiveService struct {
	PerspectiveAPIKey string
	PerspectiveUrl    string
}

func NewPerspectiveService(env *comentarylibs.Enviroment) *PerspectiveService {
	return &PerspectiveService{
		PerspectiveAPIKey: env.PerspectiveAPIKey,
		PerspectiveUrl:    env.UrlBasePerspective,
	}
}

func (s *PerspectiveService) AnalyzeText(textToAnalyze string) (*string, *bool, error) {
	body, err := getBody(textToAnalyze)
	if err != nil {
		fmt.Println("Error to get body:", err)
		return nil, nil, err
	}
	var result comentaryresponse.PerspectiveResponse
	err = jsierralibs.Post(s.PerspectiveUrl, "?key="+s.PerspectiveAPIKey, body, &result, nil)
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return nil, nil, err
	}
	score := result.AttributeScores.Toxicity.SummaryScore.Value

	isToxic := false
	if score > 0.7 {
		isToxic = true
	}

	return &textToAnalyze, &isToxic, nil
}

func getBody(textToAnalyze string) ([]byte, error) {
	// Create the request body
	requestBody := comentaryrequest.PerspectiveRequest{
		Comment: comentaryrequest.Comment{
			Text: textToAnalyze,
		},
		Languages: []string{"es"},
		RequestedAttributes: map[string]interface{}{
			"TOXICITY": map[string]interface{}{},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	return jsonData, nil
}
