package location

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/freitasmatheusrn/cloudRunLab/config"
)

type Temperature struct {
	Celsius    float64 `json:"temp_c"`
	Fahrenheit float64 `json:"temp_f"`
	Kelvin     float64 `json:"temp_k"`
}
type CEP string

func Find(cep CEP) (string, error) {
	if err := cep.Validate(); err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err

	}
	var m map[string]string
	json.Unmarshal(body, &m)
	return m["localidade"], nil
}

func Temperatures(city string) (Temperature, error) {
	cfg := config.New()
	escapedCity := url.QueryEscape(city)
	requestUrl := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", cfg.WeatherUrl, cfg.WeatherKey, escapedCity)
	response, err := http.Get(requestUrl)
	if err != nil {
		return Temperature{}, err
	}
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		errMessage := fmt.Errorf("erro ao buscar temperaturas de %s: %s", city, err)
		return Temperature{}, errMessage
	}
	var m map[string]json.RawMessage
	json.Unmarshal(body, &m)

	var current Temperature
	json.Unmarshal(m["current"], &current)
	current.Kelvin = current.Celsius + 273.15
	return current, nil
}

func (c *CEP) Validate() error {
	cep := regexp.MustCompile(`[^0-9a-zA-Z]`).ReplaceAllString(string(*c), "")
	if len(cep) != 8 {
		return errors.New("cep must be 8 numeric characters long")
	}
	if _, err := strconv.Atoi(cep); err != nil{
		return errors.New("cep cannot contain letters")

	}
	return nil
}