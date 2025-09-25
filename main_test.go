package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/freitasmatheusrn/cloudRunLab/location"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var cepFlag = flag.String("cep", "01001000", "CEP para teste de integração")

func TestIntegration_GetTemperatures_RealAPIs(t *testing.T) {
	flag.Parse()

	router := chi.NewRouter()
	router.Get("/wheater_from/{cep}", GetTemperatures)

	req := httptest.NewRequest("GET", "/wheater_from/"+*cepFlag, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "esperava 200 OK")

	var temp location.Temperature
	err := json.Unmarshal(rr.Body.Bytes(), &temp)
	temp.Kelvin = temp.Celsius + 273.15
	assert.Equal(t, temp.Kelvin, temp.Celsius+273.15)
	assert.NoError(t, err)
	output := map[string]any{
		"temperatura para o cep: " + *cepFlag: temp,
	}
	outJSON, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(outJSON))

}
