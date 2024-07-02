package articlemodel

import (
	"crud-product-golang/config"
	"crud-product-golang/entities"
)

func GetAll() []entities.Article {
	var articles []entities.Article

	rows, err := config.DB.Query(`
		SELECT 
			articles.id,
			articles.title,
			articles.body,
			categories.name as category_name,
			articles.created_at,
			articles.updated_at
		FROM articles
		JOIN categories ON articles.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article entities.Article

		err := rows.Scan(&article.Id, &article.Title, &article.Category.Name, &article.Body, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}

	return articles
}

func GetOne(id string) entities.Article {
	var article entities.Article

	err := config.DB.QueryRow(`
		SELECT 
			articles.id,
			articles.title,
			articles.body,
			categories.name as category_name,
			articles.created_at,
			articles.updated_at
		FROM articles
		JOIN categories ON articles.category_id = categories.id
		WHERE articles.id = ?
	`, id).Scan(&article.Id, &article.Title, &article.Body, &article.Category.Name, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		panic(err)
	}

	return article
}

func Create(article entities.Article) bool {
	result, err := config.DB.Exec(
		"INSERT INTO articles (title, body, category_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		article.Title, article.Body, article.Category.Id, article.CreatedAt, article.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Update(id int, article entities.Article) bool {
	result, err := config.DB.Exec(
		"UPDATE articles SET title = ?, body = ?, category_id = ?, updated_at = ? WHERE id = ?",
		article.Title, article.Body, article.Category.Id, article.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Delete(id string) bool {
	result, err := config.DB.Exec("DELETE FROM articles WHERE id = ?", id)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}