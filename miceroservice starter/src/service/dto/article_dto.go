package dto

import (
	"time"
	"web-microservice-starter/src/model"
)

// ArticleDTO - DTO of Article model
type ArticleDTO struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ArticleDTOFromArticle - Convert Article model to ArticleDTO
func ArticleDTOFromArticle(article model.Article) ArticleDTO {
	articleDTO := ArticleDTO{
		ID:        article.ID,
		Title:     article.Title,
		Body:      article.Body,
		Author:    "Gopher",
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}

	return articleDTO
}
