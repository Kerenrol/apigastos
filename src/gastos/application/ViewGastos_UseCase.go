package application

import "apiGastos/src/gastos/domain"

type ViewGastos struct {
    db domain.IGastos
}

func NewViewGastos(db domain.IGastos) *ViewGastos {
    return &ViewGastos{db: db}
}

func (vp *ViewGastos) Execute(grupoId int32) ([]domain.Gasto, error) {
    return vp.db.GetAllByGrupo(grupoId)
}