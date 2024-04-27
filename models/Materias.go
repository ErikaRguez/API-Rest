package models

import "gorm.io/gorm"

type Materias struct {
    gorm.Model
    Semestre  int `gorm:"not null" json:"Semestre"`
    Materia string `gorm:"type:varchar(255); not null" json:"Materia"`
    Creditos int `gorm:"not null" json:"Creditos"`
    CarreraID uint `json:"Id_Carrera"`
    EspecialidadID *uint  `json:"Id_Especialidad"`
    
    
}
