package main

import (

	"net/http"
	"github.com/ErikaRguez/API-Rest/db"
	"github.com/ErikaRguez/API-Rest/models"
	"github.com/ErikaRguez/API-Rest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Carreras{})
	db.DB.AutoMigrate(models.Especialidades{})
	db.DB.AutoMigrate(models.Materias{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/carrera", routes.GetCarrerasHandler).Methods("GET")
	r.HandleFunc("/carrera/{id}", routes.GetCarreraHandler).Methods("GET")
        r.HandleFunc("/carrera", routes.PostCarreraHandler).Methods("POST")
        r.HandleFunc("/carrera/{id}", routes.DeleteCarreraHandler).Methods("DELETE")

	r.HandleFunc("/especialidad", routes.GetEspecialidadesHandler).Methods("GET")
        r.HandleFunc("/especialidad/{id}", routes.GetEspecialidadHandler).Methods("GET")
        r.HandleFunc("/especialidad", routes.PostEspecialidadHandler).Methods("POST")
        r.HandleFunc("/especialidad/{id}", routes.DeleteEspecialidadHandler).Methods("DELETE")

	r.HandleFunc("/materia", routes.GetMateriasHandler).Methods("GET")
        r.HandleFunc("/materia/{id}", routes.GetMateriaHandler).Methods("GET")
        r.HandleFunc("/materia", routes.PostMateriaHandler).Methods("POST")
        r.HandleFunc("/materia/{id}", routes.DeleteMateriaHandler).Methods("DELETE")

	http.ListenAndServe(":4000", r)
}

