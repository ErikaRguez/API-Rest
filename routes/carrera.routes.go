package routes
import (
	"encoding/json"
	"net/http"
	"github.com/ErikaRguez/API-Rest/db"
	"github.com/ErikaRguez/API-Rest/models"
	"github.com/gorilla/mux"
	
)

func GetCarrerasHandler(w http.ResponseWriter, r *http.Request){
	var carreras []models.Carreras
	db.DB.Find(&carreras)
	json.NewEncoder(w).Encode(&carreras)
} 

func GetCarreraHandler(w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		var carrera models.Carreras
		db.DB.First(&carrera, params["id"])
		if carrera.ID== 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Carrera no encontrada"))
			return
		}
        db.DB.Preload("Especialidades").Preload("Materias").First(&carrera, params["id"])
		json.NewEncoder(w).Encode(&carrera)
	} 

func PostCarreraHandler(w http.ResponseWriter, r *http.Request){
	var carrera models.Carreras
	json.NewDecoder(r.Body).Decode(&carrera)
    createdCarrera := db.DB.Create(&carrera)
	err := createdCarrera.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&carrera)
} 

func DeleteCarreraHandler(w http.ResponseWriter, r *http.Request){
	var carrera models.Carreras
	params := mux.Vars(r)
	db.DB.First(&carrera, params["id"])
	if carrera.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Carrera no encontrada"))
		return
	} 
	db.DB.Unscoped().Delete(&carrera)
	w.WriteHeader(http.StatusNoContent)
} 
