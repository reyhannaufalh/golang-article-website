package categorycontroller

import (
	"crud-product-golang/entities"
	"crud-product-golang/models/categorymodel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/category/create.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	category.Name = r.FormValue("name")
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Create(category); !ok {
		temp, _ := template.ParseFiles("views/category/create.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	id := r.URL.Query().Get("id")

	category = categorymodel.GetOne(id)

	data := map[string]any{
		"category": category,
	}

	temp, err := template.ParseFiles("views/category/edit.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var category entities.Category

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	category.Name = r.FormValue("name")
	category.UpdatedAt = time.Now()

	if ok := categorymodel.Update(id, category); !ok {
		temp, _ := template.ParseFiles("views/category/edit.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	categorymodel.Delete(id)

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}