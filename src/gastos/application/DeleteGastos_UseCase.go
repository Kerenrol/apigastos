package application

import "apiGastos/src/gastos/domain"

type DeleteGasto struct {
    db domain.IGastos
}

func NewDeleteGasto(db domain.IGastos) *DeleteGasto {
    return &DeleteGasto{db: db}
}

func (dp *DeleteGasto) Execute(id int32) error {
    return dp.db.DeleteGasto(id)
}