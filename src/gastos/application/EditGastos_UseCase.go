package application

import "apiGastos/src/gastos/domain"

type EditGasto struct {
    db domain.IGastos
}

func NewEditGasto(db domain.IGastos) *EditGasto {
    return &EditGasto{db: db}
}

// Ahora recibimos descripción y monto para que la edición sea útil
func (ep *EditGasto) Execute(id int32, descripcion string, monto float64) error {
    return ep.db.UpdateGasto(id, descripcion, monto)
}