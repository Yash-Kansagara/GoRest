package server

import (
	"encoding/json"
	"net/http"

	"github.com/Yash-Kansagara/GoRest/internal/model"
)

type Product model.Product

var products []Product = []Product{
	{Id: 1, Name: "Iphone", Count: 100},
	{Id: 2, Name: "pot", Count: 10},
	{Id: 3, Name: "table", Count: 22},
}

type ProductResponse struct {
	Status int       `json:"status"`
	Count  int       `json:"count"`
	Data   []Product `json:"data"`
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	data := ProductResponse{
		Status: http.StatusOK,
		Count:  len(products),
		Data:   products,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error reading products", http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}
