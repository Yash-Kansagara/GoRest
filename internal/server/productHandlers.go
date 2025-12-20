package server

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/Yash-Kansagara/GoRest/internal/model"
)

type Product model.Product

var products []Product = []Product{
	{Id: 1, Name: "Z", Count: 100},
	{Id: 2, Name: "Y", Count: 10},
	{Id: 3, Name: "X", Count: 22},
}

var nextId uint = 4

type ProductResponse struct {
	Status int       `json:"status"`
	Count  int       `json:"count"`
	Data   []Product `json:"data"`
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	respProducts := products

	// filter
	filterID := query.Get("id")
	filterName := query.Get("name")
	respProducts = applyFilter(filterID, respProducts, filterName)

	// apply sort
	sortBy := query.Get("sortBy")
	sortOrder := query.Get("sortOrder")
	if len(sortBy) > 0 {
		applySort(sortBy, sortOrder, respProducts)
	}

	data := ProductResponse{
		Status: http.StatusOK,
		Count:  len(respProducts),
		Data:   respProducts,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error reading products", http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func applyFilter(filterID string, respProducts []Product, filterName string) []Product {
	if len(filterID) > 0 {
		idNum, e := strconv.Atoi(filterID)
		if e == nil {
			for i, p := range products {
				if p.Id == idNum {
					respProducts = respProducts[i : i+1]
					break
				}
			}
		}
	} else if len(filterName) > 0 {

		for i, p := range products {
			if p.Name == filterName {
				respProducts = respProducts[i : i+1]
				break
			}
		}
	}
	return respProducts
}

func applySort(sortBy string, sortOrder string, products []Product) {
	ascending := true
	if len(sortOrder) > 0 && sortOrder == "desc" {
		ascending = false
	}
	var sortfunc func(a Product, b Product) int
	switch sortBy {
	case "id":
		sortfunc = func(a Product, b Product) int {
			if ascending {
				return a.Id - b.Id
			}
			return b.Id - a.Id
		}

	case "name":
		sortfunc = func(a Product, b Product) int {
			if ascending {
				return strings.Compare(a.Name, b.Name)
			}
			return strings.Compare(b.Name, a.Name)
		}
	}
	slices.SortFunc(products, sortfunc)
}

func PostProductHandler(w http.ResponseWriter, r *http.Request) {

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request", http.StatusInternalServerError)
	}

	products = make([]Product, 0)
	json.Unmarshal(bodyData, &products)

	for _, p := range products {
		p.Id = int(nextId)
		nextId++
		products = append(products, p)
	}
}
