package main

import (
	"golang-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Definir rutas
	r.GET("/usuarios", controllers.GetUsuarios)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
