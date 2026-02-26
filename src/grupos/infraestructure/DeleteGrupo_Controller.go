package infraestructure

import (
	"apiGastos/src/grupos/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteGrupoController struct {
	useCase *application.DeleteGrupo
}

func NewDeleteGrupoController(useCase *application.DeleteGrupo) *DeleteGrupoController {
	return &DeleteGrupoController{useCase: useCase}
}

func (ctrl *DeleteGrupoController) Execute(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el grupo", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grupo eliminado correctamente"})
}
