package infraestructure

import (
	"apiGastos/src/gastos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteGastoController struct {
	useCase *application.DeleteGasto
}

func NewDeleteGastoController(useCase *application.DeleteGasto) *DeleteGastoController {
	return &DeleteGastoController{useCase: useCase}
}

func (ctrl *DeleteGastoController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de gasto inválido"})
		return
	}

	err = ctrl.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el gasto", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto eliminado correctamente"})
}