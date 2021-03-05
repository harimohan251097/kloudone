package controller

import (
	"net/http"
	"web-microservice-starter/src/service"
	"web-microservice-starter/src/service/dto"
	logger "web-microservice-starter/utils/logging"

	"github.com/rs/zerolog"

	"github.com/gin-gonic/gin"
)

var log *zerolog.Logger = logger.GetInstance()

// ArticleController - Article Controller
type ArticleController struct {
}

var articleService = service.ArticleService{}

// GetArticlesPath - URL Path to get Articles
const GetArticlesPath = ""

// GetArticles - List all articles
// @Summary List all articles
// @Description List all articles
// @Produce json
// @Success 200 {object} []dto.ArticleDTO "List of articles"
// @Failure 500 {object} dto.ErrorDTO "Error accessing articles"
// @Router /articles [get]
func (article *ArticleController) GetArticles(c *gin.Context) {
	limit := c.GetInt("limit")
	articleDTOs, err := articleService.GetArticles(limit)
	if err != nil {
		log.Error().Msgf("Error accessing data: %s", err.Error())
		c.JSON(http.StatusBadGateway, dto.ErrorDTO{ErrorCode: "err-500", ErrorMessage: "Error accessing articles"})
	}

	c.JSON(http.StatusOK, articleDTOs)
}

// GetArticlePath - Path to get an article
const GetArticlePath = "/:id"

// GetArticle - Get an article
// @Summary Get article by ID
// @Description Get an article by id
// @Param id path int true "id"
// @Produce json
// @Success 200 {object} dto.ArticleDTO "article"
// @Failure 500 {object} dto.ErrorDTO "Error accessing articles"
// @Router /articles/{id} [get]
func (article *ArticleController) GetArticle(c *gin.Context) {
	articleID := c.Param("id")
	articleDTO, err := articleService.GetArticle(articleID)
	if err != nil {
		log.Error().Msgf("Error accessing data: %s", err.Error())
		c.JSON(http.StatusBadGateway, dto.ErrorDTO{ErrorCode: "err-500", ErrorMessage: "Error accessing article"})
	}

	c.JSON(http.StatusOK, articleDTO)
}
