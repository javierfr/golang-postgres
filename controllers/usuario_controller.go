package controllers

import (
	"context"
	"net/http"

	"golang-postgres/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetUsuarios(c *gin.Context) {
	// URL de conexión a PostgreSQL
	// urlConn := "postgres://postgres:root@localhost:5432/db_golang"
	urlConn := "postgres://postgres:root@localhost:5432/db_golang"

	// Conectar a la base de datos
	conn, err := pgx.Connect(context.Background(), urlConn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	defer conn.Close(context.Background())

	usuarios, err := models.GetAllUsuarios(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}
