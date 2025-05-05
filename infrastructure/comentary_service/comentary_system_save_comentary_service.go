package comentaryservice

import (
	"context"
	"time"

	comentarymodels "github.com/jSierraB3991/comentary_system/domain/comentary_models"
	cometaryvalidations "github.com/jSierraB3991/comentary_system/domain/cometary_validations"
)

func (s *ComentarySystemService) SaveComentary(ctx context.Context, comentary, commentUpId string, idPost, idUserComment uint) error {
	err := cometaryvalidations.ValidateComentary(comentary)
	if err != nil {
		return err
	}

	isToxic, err := s.perspectiveService.AnalyzeText(comentary)
	if err != nil {
		return err
	}

	comentaryModel := comentarymodels.ComentarySystemComentaryModel{
		IdUser:    idUserComment,
		IsToxic:   *isToxic,
		IdPost:    idPost,
		Comentary: comentary,
		CreatedAt: time.Now().UTC(),
	}
	return s.repository.SaveComentary(ctx, comentaryModel, commentUpId)
}

func (s *ComentarySystemService) GetComentarysByUser(ctx context.Context, userId uint) ([]comentarymodels.ComentarySystemComentaryModel, error) {
	comentarys, err := s.repository.GetComentarysByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return comentarys, nil
}
