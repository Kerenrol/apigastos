package infraestructure

import (
	"apiGastos/src/grupos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetGrupoByIdController struct {
	useCase *application.GetGrupoById
}

func NewGetGrupoByIdController(useCase *application.GetGrupoById) *GetGrupoByIdController {
	return &GetGrupoByIdController{useCase: useCase}
}

func (ctrl *GetGrupoByIdController) Execute(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	grupo, err := ctrl.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el grupo", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grupo)
}
