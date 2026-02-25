package infraestructure

import (
    "apiGastos/src/gastos/application"
    "apiGastos/src/gastos/domain"

    "github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IGastos, r *gin.Engine) {
    createGastoUseCase := application.NewCreateGasto(repo)
    createGastoController := NewCreateGastoController(createGastoUseCase)

    viewGastosUseCase := application.NewViewGastos(repo)
    viewGastosController := NewViewGastosController(viewGastosUseCase)

    editGastoUseCase := application.NewEditGasto(repo)
    editGastoController := NewEditGastoController(editGastoUseCase)

    deleteGastoUseCase := application.NewDeleteGasto(repo)
    deleteGastoController := NewDeleteGastoController(deleteGastoUseCase)

    getGastoByIdUseCase := application.NewGetGastoById(repo)
    getGastoByIdController := NewGetGastoByIdController(getGastoByIdUseCase)
    
    api := r.Group("/gastos")
    {
        api.POST("/", createGastoController.Execute)        // Crear un nuevo gasto
        api.GET("/", viewGastosController.Execute)         // Listar todos los gastos
        api.GET("/:id", getGastoByIdController.Execute)    // Detalle de un gasto específico
        api.PUT("/:id", editGastoController.Execute)       // Editar un gasto existente
        api.DELETE("/:id", deleteGastoController.Execute)  // Eliminar un gasto
    }
}