package service

import (
	"web-microservice-starter/src/repository"
	"web-microservice-starter/src/service/dto"
)

// ArticleService -
type ArticleService struct{}

var articleRepository repository.ArticleRepository = repository.ArticleRepository{}

// GetArticles - Get articles
func (articleService *ArticleService) GetArticles(limit int) ([]dto.ArticleDTO, error) {
	articles, err := articleRepository.GetAllArticles(limit)
	if err != nil {
		return nil, err
	}

	articleDTOs := make([]dto.ArticleDTO, len(articles))
	for _, article := range articles {
		articleDTOs = append(articleDTOs, dto.ArticleDTOFromArticle(article))
	}
	return articleDTOs, nil
}

func (articleService *ArticleService) GetArticle(articleId string) (*dto.ArticleDTO, error) {
	article, err := articleRepository.GetArticle(articleId)
	if err != nil {
		return nil, err
	}

	articleDTO := dto.ArticleDTOFromArticle(*article)
	return &articleDTO, nil
}
