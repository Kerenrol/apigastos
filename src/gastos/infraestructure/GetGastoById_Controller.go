package infraestructure

import (
	"apiGastos/src/gastos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetGastoByIdController struct {
	useCase *application.GetGastoById
}

func NewGetGastoByIdController(useCase *application.GetGastoById) *GetGastoByIdController {
	return &GetGastoByIdController{useCase: useCase}
}

func (ctrl *GetGastoByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de gasto inválido"})
		return
	}

	gasto, err := ctrl.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el gasto", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gasto)
}