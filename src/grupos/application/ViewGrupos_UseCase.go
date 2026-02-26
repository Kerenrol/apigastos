package application

import "apiGastos/src/grupos/domain"

type ViewGrupos struct {
	db domain.IGrupos
}

func NewViewGrupos(db domain.IGrupos) *ViewGrupos {
	return &ViewGrupos{db: db}
}

func (vg *ViewGrupos) Execute() ([]domain.Grupo, error) {
	return vg.db.GetAllGrupos()
}
