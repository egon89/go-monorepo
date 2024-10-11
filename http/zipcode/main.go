package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func main() {
	fmt.Println("Server is listening on port 8080...")
	http.HandleFunc("/", GetZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func GetZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipCodeParam := r.URL.Query().Get("zip-code")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zipCode, err := GetZipCode(zipCodeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipCode)
}

func GetZipCode(zipCode string) (*ViaCEP, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipCode + "/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var zc ViaCEP
	err = json.Unmarshal(body, &zc)
	if err != nil {
		return nil, err
	}
	return &zc, nil
}
