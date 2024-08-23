package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Correo        string `json:"correo"`
	Contrasena    string `json:"contrasena"`
	FechaCreacion string `json:"fecha_creacion"`
	UltimoAcceso  string `json:"ultimo_acceso"`
}

func GetAllUsuarios(conn *pgx.Conn) ([]Usuario, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, nombre, correo, contrasena, fecha_creacion, ultimo_acceso FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Correo, &usuario.Contrasena, &usuario.FechaCreacion, &usuario.UltimoAcceso)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if rows.Err() != nil {
		log.Printf("Error al iterar las filas: %v\n", rows.Err())
		return nil, rows.Err()
	}

	return usuarios, nil
}
