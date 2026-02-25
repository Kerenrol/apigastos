package infraestructure

import (
    "apiGastos/src/gastos/application"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type ViewGastosController struct {
    useCase *application.ViewGastos
}

func NewViewGastosController(useCase *application.ViewGastos) *ViewGastosController {
    return &ViewGastosController{useCase: useCase}
}

func (ctrl *ViewGastosController) Execute(c *gin.Context) {
    // Obtenemos el grupo_id de la URL ?grupo_id=1
    grupoIdStr := c.Query("grupo_id")
    grupoId, _ := strconv.Atoi(grupoIdStr)

    gastos, err := ctrl.useCase.Execute(int32(grupoId))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar los gastos", "detalles": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"gastos": gastos})
}