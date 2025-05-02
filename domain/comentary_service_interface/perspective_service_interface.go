package comentaryserviceinterface

type PerspectiveApiServiceInterface interface {
	AnalyzeText(textToAnalyze string) (*string, *bool, error)
}
