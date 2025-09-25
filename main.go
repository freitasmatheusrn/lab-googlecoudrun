package main

import (
	"encoding/json"
	"net/http"

	"github.com/freitasmatheusrn/cloudRunLab/location"
	"github.com/go-chi/chi/v5"
)

func GetTemperatures(w http.ResponseWriter, r *http.Request) {
	cep := location.CEP(chi.URLParam(r, "cep"))
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cep is empty"))
		return
	}
	err := cep.Validate()
	if err != nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("invalid zip code:" + err.Error()))
		return
	}
	city, err := location.Find(cep)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	temperatures, err := location.Temperatures(city)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	output := map[string]any{
		"temperatura para a cidade de " + city: temperatures,
	}
	
	json.NewEncoder(w).Encode(output)
}

func main() {
	router := chi.NewRouter()
	router.Get("/wheater_from/{cep}", GetTemperatures)
	http.ListenAndServe(":8080", router)
}
