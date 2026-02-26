package infraestructure

import (
	"apiGastos/src/gastos/application"
	"apiGastos/src/gastos/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGastoController struct {
	useCase *application.CreateGasto
	repo    domain.IGastos
}

func NewCreateGastoController(useCase *application.CreateGasto, repo domain.IGastos) *CreateGastoController {
	return &CreateGastoController{useCase: useCase, repo: repo}
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

	id, err := ctrl.useCase.Execute(body.Descripcion, body.Monto, body.PagadorID, body.GrupoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el gasto", "detalles": err.Error()})
		return
	}

	gasto, err := ctrl.repo.GetGastoById(id)
	if err == nil {
		GetHub().BroadcastEvent("create", body.GrupoID, gasto)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Gasto registrado correctamente"})
}
