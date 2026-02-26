package application

import "apiGastos/src/grupos/domain"

type DeleteGrupo struct {
	db domain.IGrupos
}

func NewDeleteGrupo(db domain.IGrupos) *DeleteGrupo {
	return &DeleteGrupo{db: db}
}

func (dg *DeleteGrupo) Execute(id int32) error {
	return dg.db.DeleteGrupo(id)
}
