package repository

import (
	"strconv"
	"web-microservice-starter/src/model"
	"web-microservice-starter/utils/database"
	logger "web-microservice-starter/utils/logging"

	"github.com/rs/zerolog"
)

// ArticleRepository -
type ArticleRepository struct{}

var log *zerolog.Logger = logger.GetInstance()

// GetAllArticles -
func (articleRepository ArticleRepository) GetAllArticles(limit int) ([]model.Article, error) {
	var articles []model.Article

	if err := database.GetInstance().Find(&articles).Limit(limit).Error; err != nil {
		return nil, err
	}

	log.Info().Msgf("Articles %v", articles)
	return articles, nil
}

func (articleRepository *ArticleRepository) GetArticle(articleId string) (*model.Article, error) {
	var article model.Article
	// articleID := strconv.Atoi(articleId)
	if articleID, err := strconv.Atoi(articleId); err == nil {
		article.ID = articleID
	}

	if err := database.GetInstance().Find(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}
