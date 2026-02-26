package infraestructure

import (
	"apiGastos/src/gastos/application"
	"apiGastos/src/gastos/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditGastoController struct {
	useCase *application.EditGasto
	repo    domain.IGastos
}

func NewEditGastoController(useCase *application.EditGasto, repo domain.IGastos) *EditGastoController {
	return &EditGastoController{useCase: useCase, repo: repo}
}

func (ctrl *EditGastoController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de gasto inválido"})
		return
	}

	var body struct {
		Descripcion string  `json:"descripcion"`
		Monto       float64 `json:"monto"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos de actualización"})
		return
	}

	// Obtener el gasto original para saber su grupo_id
	gastoOriginal, err := ctrl.repo.GetGastoById(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el gasto", "detalles": err.Error()})
		return
	}

	err = ctrl.useCase.Execute(int32(id), body.Descripcion, body.Monto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el gasto", "detalles": err.Error()})
		return
	}

	// Emitir evento a través de WebSocket
	GetHub().BroadcastEvent("update", gastoOriginal.GrupoID, gin.H{
		"id":          id,
		"descripcion": body.Descripcion,
		"monto":       body.Monto,
		"grupo_id":    gastoOriginal.GrupoID,
		"pagador_id":  gastoOriginal.PagadorID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Gasto actualizado correctamente"})
}
