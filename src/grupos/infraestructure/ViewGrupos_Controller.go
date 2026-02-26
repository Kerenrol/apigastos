package infraestructure

import (
	"apiGastos/src/grupos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewGruposController struct {
	useCase *application.ViewGrupos
}

func NewViewGruposController(useCase *application.ViewGrupos) *ViewGruposController {
	return &ViewGruposController{useCase: useCase}
}

func (ctrl *ViewGruposController) Execute(c *gin.Context) {
	grupos, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener grupos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grupos)
}
