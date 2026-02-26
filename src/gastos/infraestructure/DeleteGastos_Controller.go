package infraestructure

import (
	"apiGastos/src/gastos/application"
	"apiGastos/src/gastos/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteGastoController struct {
	useCase *application.DeleteGasto
	repo    domain.IGastos
}

func NewDeleteGastoController(useCase *application.DeleteGasto, repo domain.IGastos) *DeleteGastoController {
	return &DeleteGastoController{useCase: useCase, repo: repo}
}

func (ctrl *DeleteGastoController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de gasto inválido"})
		return
	}

	// Obtener el gasto original para saber su grupo_id
	gastoOriginal, err := ctrl.repo.GetGastoById(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el gasto", "detalles": err.Error()})
		return
	}

	err = ctrl.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el gasto", "detalles": err.Error()})
		return
	}

	// Emitir evento a través de WebSocket
	GetHub().BroadcastEvent("delete", gastoOriginal.GrupoID, gin.H{
		"id":       id,
		"grupo_id": gastoOriginal.GrupoID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Gasto eliminado correctamente"})
}
