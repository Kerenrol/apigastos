package infraestructure

import (
	"apiGastos/src/grupos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGrupoController struct {
	useCase *application.CreateGrupo
}

func NewCreateGrupoController(useCase *application.CreateGrupo) *CreateGrupoController {
	return &CreateGrupoController{useCase: useCase}
}

type RequestBodyGrupo struct {
	Nombre string `json:"nombre"`
}

func (ctrl *CreateGrupoController) Execute(c *gin.Context) {
	var body RequestBodyGrupo
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos del grupo inválidos", "detalles": err.Error()})
		return
	}

	err := ctrl.useCase.Execute(body.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el grupo", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Grupo registrado correctamente"})
}
