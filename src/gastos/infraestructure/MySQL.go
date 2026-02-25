package infraestructure

import (
	"fmt"
	"log"

	"apiGastos/src/config"
	"apiGastos/src/gastos/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IGastos = (*MySQL)(nil)

func NewMySQL() domain.IGastos {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateGasto(gasto *domain.Gasto) error {
	query := "INSERT INTO gastos (descripcion, monto, pagador_id, grupo_id, fecha) VALUES (?, ?, ?, ?, NOW())"
	_, err := mysql.conn.ExecutePreparedQuery(query, gasto.Descripcion, gasto.Monto, gasto.PagadorID, gasto.GrupoID)
	if err != nil {
		return fmt.Errorf("error al insertar gasto: %v", err)
	}
	return nil
}

func (mysql *MySQL) GetGastoById(id int32) (*domain.Gasto, error) {
	query := "SELECT id, descripcion, monto, pagador_id, grupo_id, fecha FROM gastos WHERE id = ?"
	row, err := mysql.conn.FetchRow(query, id)
	if err != nil {
		return nil, err
	}

	var gasto domain.Gasto
	if err := row.Scan(&gasto.ID, &gasto.Descripcion, &gasto.Monto, &gasto.PagadorID, &gasto.GrupoID, &gasto.Fecha); err != nil {
		return nil, err
	}
	return &gasto, nil
}

func (mysql *MySQL) GetAllByGrupo(grupoId int32) ([]domain.Gasto, error) {
	query := "SELECT id, descripcion, monto, pagador_id, grupo_id, fecha FROM gastos WHERE grupo_id = ?"
	rows, err := mysql.conn.FetchRows(query, grupoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gastos []domain.Gasto
	for rows.Next() {
		var g domain.Gasto
		rows.Scan(&g.ID, &g.Descripcion, &g.Monto, &g.PagadorID, &g.GrupoID, &g.Fecha)
		gastos = append(gastos, g)
	}
	return gastos, nil
}

func (mysql *MySQL) UpdateGasto(id int32, descripcion string, monto float64) error {
	query := "UPDATE gastos SET descripcion = ?, monto = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, descripcion, monto, id)
	return err
}

func (mysql *MySQL) DeleteGasto(id int32) error {
	query := "DELETE FROM gastos WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	return err
}

func (mysql *MySQL) GetSaldos(grupoId int32) (map[int32]float64, error) {
	// Esta consulta calcula cuánto ha pagado cada usuario vs el promedio del grupo
	query := `
		SELECT pagador_id, 
		SUM(monto) - (SELECT SUM(monto)/COUNT(DISTINCT pagador_id) FROM gastos WHERE grupo_id = ?) as saldo
		FROM gastos 
		WHERE grupo_id = ?
		GROUP BY pagador_id`
	
	rows, err := mysql.conn.FetchRows(query, grupoId, grupoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	saldos := make(map[int32]float64)
	for rows.Next() {
		var userID int32
		var saldo float64
		if err := rows.Scan(&userID, &saldo); err != nil {
			return nil, err
		}
		saldos[userID] = saldo
	}
	return saldos, nil
}