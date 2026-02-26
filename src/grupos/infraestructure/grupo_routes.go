package infraestructure

import (
	"apiGastos/src/grupos/application"
	"apiGastos/src/grupos/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IGrupos, r *gin.Engine) {
	createGrupoUseCase := application.NewCreateGrupo(repo)
	createGrupoController := NewCreateGrupoController(createGrupoUseCase)

	viewGruposUseCase := application.NewViewGrupos(repo)
	viewGruposController := NewViewGruposController(viewGruposUseCase)

	getGrupoByIdUseCase := application.NewGetGrupoById(repo)
	getGrupoByIdController := NewGetGrupoByIdController(getGrupoByIdUseCase)

	deleteGrupoUseCase := application.NewDeleteGrupo(repo)
	deleteGrupoController := NewDeleteGrupoController(deleteGrupoUseCase)

	api := r.Group("/grupos")
	{
		api.POST("/", createGrupoController.Execute)      // Crear un nuevo grupo
		api.GET("/", viewGruposController.Execute)        // Listar todos los grupos
		api.GET("/:id", getGrupoByIdController.Execute)   // Detalle de un grupo específico
		api.DELETE("/:id", deleteGrupoController.Execute) // Eliminar un grupo
	}
}

func Init(r *gin.Engine) {
	repo := NewMySQL()
	SetupRouter(repo, r)
}
