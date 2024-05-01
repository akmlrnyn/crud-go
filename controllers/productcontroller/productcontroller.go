package productcontroller

import (
	"crud-go/entities"
	"crud-go/models/categorymodel"
	"crud-go/models/productmodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products" : products,
	}

	temp, err := template.ParseFiles("views/product/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")

		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		id, err := strconv.Atoi(r.FormValue("category_id"))

		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))

		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(id)
		product.Stock =  int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()
		product.UpdatedAt = time.Now()
		
		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}