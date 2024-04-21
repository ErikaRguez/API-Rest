package models

import "gorm.io/gorm"

type Especialidades struct {
    gorm.Model
    Especialidad string  `gorm:"type:varchar(255); not null" json:"Especialidad"`
    CarreraID  uint `json:"Id_Carrera"`
    Materias  []Materias `gorm:"foreignKey:CarreraID" json:"Materias"`
}
