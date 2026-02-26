package domain

type Gasto struct {
	ID          int32   `json:"id"`
	Descripcion string  `json:"descripcion"`
	Monto       float64 `json:"monto"`
	PagadorID   int32   `json:"pagador_id"`
	GrupoID     int32   `json:"grupo_id"` // Para saber en qué viaje/grupo se hizo
	Fecha       string  `json:"fecha"`
}

type IGastos interface {
	CreateGasto(gasto *Gasto) (int32, error)
	GetAllByGrupo(grupoId int32) ([]Gasto, error)
	GetGastoById(id int32) (*Gasto, error)
	UpdateGasto(id int32, descripcion string, monto float64) error
	DeleteGasto(id int32) error
	GetSaldos(grupoId int32) (map[int32]float64, error)
}
