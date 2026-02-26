package application

import "apiGastos/src/grupos/domain"

type GetGrupoById struct {
	db domain.IGrupos
}

func NewGetGrupoById(db domain.IGrupos) *GetGrupoById {
	return &GetGrupoById{db: db}
}

func (gg *GetGrupoById) Execute(id int32) (*domain.Grupo, error) {
	return gg.db.GetGrupoById(id)
}
