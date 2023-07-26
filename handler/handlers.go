package handler

import (
	"go-web/entity"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) { //Handler
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Name":    "Fahmi",
		"Message": "Belajar Go Lang",
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		http.Error(w, "Error is happening because file is not found", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error is happening because file is not found", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("Home1"))
}
func HelloFahmi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Fahmi"))
}

func Anonymouse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SPY VS SPY"))
}

func Product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 0 {
		http.NotFound(w, r)
		return
	}
	data := []entity.Product{
		{
			ID:    idNumb,
			Name:  "Product 1",
			Price: 1000,
			Stock: 10,
		},
		{
			ID:    idNumb + 1,
			Name:  "Product 2",
			Price: 2000,
			Stock: 20,
		},
		{
			ID:    idNumb + 2,
			Name:  "Product 3",
			Price: 3000,
			Stock: 30,
		},
	}
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"))

	if err != nil {
		http.NotFound(w, r)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}
