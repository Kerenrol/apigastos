package main

import (
	"log"

	gastosInfra "apiGastos/src/gastos/infraestructure"
	gruposInfra "apiGastos/src/grupos/infraestructure"
	usersInfra "apiGastos/src/users/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	usersInfra.InitRouter(r, usersInfra.NewMySQL())
	gruposInfra.Init(r)
	gastosInfra.Init(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
