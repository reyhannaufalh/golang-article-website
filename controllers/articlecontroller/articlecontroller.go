package articlecontroller

import (
	"crud-product-golang/entities"
	"crud-product-golang/models/articlemodel"
	"crud-product-golang/models/categorymodel"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	articles := articlemodel.GetAll()

	data := map[string]any{
		"articles": articles,
	}

	temp, err := template.ParseFiles("views/article/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/article/create.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var article entities.Article

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))

	if err != nil {
		panic(err)
	}

	article.Title = r.FormValue("title")
	article.Body = r.FormValue("body")
	article.Category.Id = uint(categoryId)
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	log.Println(article)

	if ok := articlemodel.Create(article); !ok {
		temp, _ := template.ParseFiles("views/article/create.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var article entities.Article

	id := r.URL.Query().Get("id")

	article = articlemodel.GetOne(id)
	categories := categorymodel.GetAll()

	data := map[string]any{
		"article":    article,
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/article/edit.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var article entities.Article

	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(r.FormValue("category_id"))

	if err != nil {
		panic(err)
	}

	article.Id = uint(id)
	article.Title = r.FormValue("title")
	article.Body = r.FormValue("body")
	article.Category.Id = uint(categoryId)
	article.UpdatedAt = time.Now()

	if ok := articlemodel.Update(id, article); !ok {
		temp, _ := template.ParseFiles("views/article/edit.html")
		temp.Execute(w, nil)
	}

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	articlemodel.Delete(id)

	http.Redirect(w, r, "/articles", http.StatusSeeOther)
}