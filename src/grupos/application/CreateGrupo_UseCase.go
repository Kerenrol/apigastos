package application

import "apiGastos/src/grupos/domain"

type CreateGrupo struct {
	db domain.IGrupos
}

func NewCreateGrupo(db domain.IGrupos) *CreateGrupo {
	return &CreateGrupo{db: db}
}

func (cg *CreateGrupo) Execute(nombre string) error {
	nuevoGrupo := &domain.Grupo{
		Nombre: nombre,
	}
	return cg.db.CreateGrupo(nuevoGrupo)
}
