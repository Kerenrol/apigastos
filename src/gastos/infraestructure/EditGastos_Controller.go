package infraestructure

import (
	"apiGastos/src/gastos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditGastoController struct {
	useCase *application.EditGasto
}

func NewEditGastoController(useCase *application.EditGasto) *EditGastoController {
	return &EditGastoController{useCase: useCase}
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

	err = ctrl.useCase.Execute(int32(id), body.Descripcion, body.Monto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el gasto", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto actualizado correctamente"})
}