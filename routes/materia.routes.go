package routes

import (
	"encoding/json"
	"net/http"
	"github.com/ErikaRguez/API-Rest/db"
	"github.com/ErikaRguez/API-Rest/models"
	"github.com/gorilla/mux"	
)

func GetMateriasHandler(w http.ResponseWriter, r *http.Request){
	var materias []models.Materias
	db.DB.Find(&materias)
	json.NewEncoder(w).Encode(materias)
}

func GetMateriaHandler(w http.ResponseWriter, r *http.Request){
	var materia models.Materias
	params := mux.Vars(r)
	db.DB.First(&materia, params["id"])
	if materia.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Materia no encontrada"))
		return
	}
	json.NewEncoder(w).Encode(&materia)
}

func PostMateriaHandler(w http.ResponseWriter, r *http.Request){
	var materia models.Materias
	json.NewDecoder(r.Body).Decode(&materia)
	db.DB.Create(&materia)
	createdMateria := db.DB.Create(&materia)
	err := createdMateria.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&materia)
}

func DeleteMateriaHandler(w http.ResponseWriter, r *http.Request){
	var materia models.Materias
	params := mux.Vars(r)
	db.DB.First(&materia, params["id"])
	if materia.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Materia no encontrada"))
		return
	}
	db.DB.Unscoped().Delete(&materia)
	w.WriteHeader(http.StatusNoContent)
}
	
