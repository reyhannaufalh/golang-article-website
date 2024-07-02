package entities

import "time"

type Article struct {
	Id        uint
	Title     string
	Body      string
	Category  Category
	CreatedAt time.Time
	UpdatedAt time.Time
}