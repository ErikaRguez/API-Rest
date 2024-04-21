package models

import "gorm.io/gorm"

type Carreras struct {
    gorm.Model
    Carrera  string `gorm:"type:varchar(255); not null" json:"Carrera"`
    Especialidades []Especialidades `gorm:"foreignKey:CarreraID" json:"Especialidades"`
    Materias  []Materias `gorm:"foreignKey:CarreraID" json:"Materias"`
}
