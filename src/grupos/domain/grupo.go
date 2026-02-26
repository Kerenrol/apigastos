package domain

type Grupo struct {
	ID       int32  `json:"id"`
	Nombre   string `json:"nombre"`
	CreadoEn string `json:"creado_en"`
}

type IGrupos interface {
	CreateGrupo(grupo *Grupo) error
	GetAllGrupos() ([]Grupo, error)
	GetGrupoById(id int32) (*Grupo, error)
	DeleteGrupo(id int32) error
}
