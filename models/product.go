package models

import "gorm.io/gorm"

type Producto struct {
	gorm.Model
	ID       int64   `json:"id" gorm:"primary_key"`
	Nombre   string  `json:"nombre"`
	Cantidad int64   `json:"cantidad"`
	Precio   float64 `json:"precio"`
}
