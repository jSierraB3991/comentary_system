package comentaryserviceinterface

type PerspectiveApiServiceInterface interface {
	AnalyzeText(textToAnalyze string) (*bool, error)
}
