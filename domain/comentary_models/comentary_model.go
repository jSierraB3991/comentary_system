package comentarymodels

import "time"

type ComentarySystemComentaryModel struct {
	Id        string                          `json:"id"`
	IdPost    uint                            `json:"id_post"`
	IdUser    uint                            `json:"id_user"`
	IsToxic   bool                            `json:"is_toxic"`
	Comentary string                          `json:"comentary"`
	CreatedAt time.Time                       `json:"created_at"`
	Coments   []ComentarySystemComentaryModel `json:"coments"`
}
