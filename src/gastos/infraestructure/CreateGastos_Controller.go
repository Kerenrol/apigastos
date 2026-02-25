package infraestructure

import (
	"apiGastos/src/gastos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGastoController struct {
	useCase *application.CreateGasto
}

func NewCreateGastoController(useCase *application.CreateGasto) *CreateGastoController {
	return &CreateGastoController{useCase: useCase}
}

type RequestBody struct {
	Descripcion string  `json:"descripcion"`
	Monto       float64 `json:"monto"`
	PagadorID   int32   `json:"pagador_id"`
	GrupoID     int32   `json:"grupo_id"`
}

func (ctrl *CreateGastoController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos del gasto inválidos", "detalles": err.Error()})
		return
	}

	err := ctrl.useCase.Execute(body.Descripcion, body.Monto, body.PagadorID, body.GrupoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el gasto", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gasto registrado correctamente"})
}