package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Yash-Kansagara/GoRest/internal/db"
	"github.com/Yash-Kansagara/GoRest/internal/model"
)

type Product model.Product

type ProductResponse struct {
	Status int       `json:"status"`
	Count  int       `json:"count"`
	Data   []Product `json:"data"`
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	logReqDetails(r)
	switch r.Method {
	case http.MethodGet:
		GetProductHandler(w, r)
	case http.MethodPost:
		PostProductHandler(w, r)
	case http.MethodPatch:
		w.Write([]byte("products PATCH"))
	case http.MethodPut:
		w.Write([]byte("products PUT"))
	case http.MethodDelete:
		w.Write([]byte("products DELETE"))
	}
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {

	var stringBuilder *strings.Builder = &strings.Builder{}
	query := r.URL.Query()

	stringBuilder.WriteString("SELECT * FROM products where 1=1 ")
	// filter
	filterID := query.Get("id")
	filterName := query.Get("name")
	applyFilter(stringBuilder, filterID, filterName)

	// apply sort
	sortBy := query.Get("sortBy")
	sortOrder := query.Get("sortOrder")
	applySort(stringBuilder, sortBy, sortOrder)

	stringBuilder.WriteRune(';')
	// fetch data from db
	db := db.GetDB()
	stmt, err := db.Prepare(stringBuilder.String())
	if err != nil {
		log.Println(err)
		http.Error(w, "ERROR fetching data 101", http.StatusInternalServerError)
		return
	}

	rows, err := stmt.Query()

	if err != nil {
		http.Error(w, "ERROR fetching data 102", http.StatusInternalServerError)
	}

	p := Product{}
	resp := []Product{}
	for rows.Next() {
		rows.Scan(&p.Id, &p.Name, &p.Count)
		resp = append(resp, p)
	}
	data := ProductResponse{
		Status: http.StatusOK,
		Count:  len(resp),
		Data:   resp,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error reading products", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func applyFilter(sb *strings.Builder, filterID string, filterName string) *strings.Builder {
	if len(filterID) > 0 {
		sb.WriteString(fmt.Sprintf("AND id='%s' ", filterID))
	}
	if len(filterName) > 0 {

		sb.WriteString(fmt.Sprintf("AND name='%s' ", filterName))
	}
	return sb
}

func applySort(sb *strings.Builder, sortBy string, sortOrder string) {
	if len(sortBy) > 0 {
		sortOrderQuery := "ASC"
		if len(sortOrder) > 0 && sortOrder == "desc" {
			sortOrderQuery = "DESC"
		}

		sb.WriteString(fmt.Sprintf("ORDER BY %s %s", sortBy, sortOrderQuery))
	}
}

func PostProductHandler(w http.ResponseWriter, r *http.Request) {

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request", http.StatusInternalServerError)
	}

	products := make([]Product, 0)
	json.Unmarshal(bodyData, &products)

	db := db.GetDB()
	queryBuilder := "INSERT INTO products (name, count) VALUES(?,?)"
	stmt, err := db.Prepare(queryBuilder)
	if err != nil {
		http.Error(w, "ERROR Adding Product ", http.StatusInternalServerError)
		return
	}

	for i, p := range products {
		res, err := stmt.Exec(p.Name, p.Count)
		if err != nil {
			log.Println("failed to insert", p)
			continue
		}

		id, _ := res.LastInsertId()
		products[i].Id = int(id)
	}

	bodyData, err = json.Marshal(products)
	if err != nil {
		http.Error(w, "ERROR parsing Products", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(bodyData)
}
