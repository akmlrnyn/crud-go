package productcontroller

import (
	"crud-go/models/productmodel"
	"net/http"
	"text/template"
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