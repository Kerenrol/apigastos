package application

import "apiGastos/src/gastos/domain"

type GetGastoById struct {
    db domain.IGastos
}

func NewGetGastoById(db domain.IGastos) *GetGastoById {
    return &GetGastoById{db: db}
}

func (gp *GetGastoById) Execute(id int32) (*domain.Gasto, error) {
    return gp.db.GetGastoById(id)
}