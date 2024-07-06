package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model        // agrega por defecto los campos ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func (Person) TableName() string {
	return "persons" // nombre de la tabla, por defecto gorm toma el nombre de la clase en plural y minuscula
}
