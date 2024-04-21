package routes

import (
	"encoding/json"
	"net/http"
	"github.com/ErikaRguez/API-Rest/db"
	"github.com/ErikaRguez/API-Rest/models"
	"github.com/gorilla/mux"
	
)

func GetEspecialidadesHandler(w http.ResponseWriter, r *http.Request){
	var especialidades []models.Especialidades
	db.DB.Find(&especialidades)
	json.NewEncoder(w).Encode(especialidades)

}

func GetEspecialidadHandler(w http.ResponseWriter, r *http.Request){
	var especialidad models.Especialidades
	params := mux.Vars(r)
	db.DB.First(&especialidad, params["id"])
	if especialidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Especialidad no encontrada"))
		return
	}
	db.DB.Preload("Materias").First(&especialidad, params["id"])
	json.NewEncoder(w).Encode(&especialidad)
}

func PostEspecialidadHandler(w http.ResponseWriter, r *http.Request){
	var especialidad models.Especialidades
	json.NewDecoder(r.Body).Decode(&especialidad)
	db.DB.Create(&especialidad)
	createdEspecialidad := db.DB.Create(&especialidad)
	err := createdEspecialidad.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&especialidad)
}

func DeleteEspecialidadHandler(w http.ResponseWriter, r *http.Request){
	var especialidad models.Especialidades
	params := mux.Vars(r)
	db.DB.First(&especialidad, params["id"])
	if especialidad.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Especialidad no encontrada"))
		return
	}
	db.DB.Unscoped().Delete(&especialidad)
	w.WriteHeader(http.StatusNoContent)
}
	
