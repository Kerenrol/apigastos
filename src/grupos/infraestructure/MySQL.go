package infraestructure

import (
	"fmt"
	"log"

	"apiGastos/src/config"
	"apiGastos/src/grupos/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IGrupos = (*MySQL)(nil)

func NewMySQL() domain.IGrupos {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateGrupo(grupo *domain.Grupo) error {
	query := "INSERT INTO grupos (nombre) VALUES (?)"
	_, err := mysql.conn.ExecutePreparedQuery(query, grupo.Nombre)
	if err != nil {
		return fmt.Errorf("error al insertar grupo: %v", err)
	}
	return nil
}

func (mysql *MySQL) GetAllGrupos() ([]domain.Grupo, error) {
	query := "SELECT id, nombre, creado_en FROM grupos"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grupos []domain.Grupo
	for rows.Next() {
		var g domain.Grupo
		if err := rows.Scan(&g.ID, &g.Nombre, &g.CreadoEn); err != nil {
			return nil, err
		}
		grupos = append(grupos, g)
	}
	return grupos, nil
}

func (mysql *MySQL) GetGrupoById(id int32) (*domain.Grupo, error) {
	query := "SELECT id, nombre, creado_en FROM grupos WHERE id = ?"
	row, err := mysql.conn.FetchRow(query, id)
	if err != nil {
		return nil, err
	}

	var grupo domain.Grupo
	if err := row.Scan(&grupo.ID, &grupo.Nombre, &grupo.CreadoEn); err != nil {
		return nil, err
	}
	return &grupo, nil
}

func (mysql *MySQL) DeleteGrupo(id int32) error {
	query := "DELETE FROM grupos WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	return err
}
