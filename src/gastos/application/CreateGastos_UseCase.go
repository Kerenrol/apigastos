package application

import "apiGastos/src/gastos/domain"

type CreateGasto struct {
	db domain.IGastos
}

func NewCreateGasto(db domain.IGastos) *CreateGasto {
	return &CreateGasto{db: db}
}

func (cg *CreateGasto) Execute(desc string, monto float64, pagadorId int32, grupoId int32) (int32, error) {
	nuevoGasto := &domain.Gasto{
		Descripcion: desc,
		Monto:       monto,
		PagadorID:   pagadorId,
		GrupoID:     grupoId,
	}
	return cg.db.CreateGasto(nuevoGasto)
}
