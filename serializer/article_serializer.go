package serializer

import (
	"gin-blog/models"
)

// ArticleSerializer 文章序列化器
type ArticleSerializer struct {
	Id        uint64             `json:"id"`
	Title     string             `json:"title"`
	Author    string             `json:"author"`
	Cover     string             `json:"cover"`
	Content   string             `json:"content"`
	Status    int                `json:"status"`
	IsHot     int                `json:"is_hot"`
	ReadCount int                `json:"read_count"`
	Cid       int                `json:"cid"`
	Category  CategorySerializer `json:"category"`
	CreatedAt int64              `json:"created_at"`
	UpdatedAt int64              `json:"updated_at"`
}

// BuildArticle 序列化文章
func BuildArticle(article *models.Article) ArticleSerializer {
	return ArticleSerializer{
		Id:        article.Id,
		Title:     article.Title,
		Author:    article.Author,
		Cover:     article.Cover,
		Content:   article.Content,
		Status:    article.Status,
		IsHot:     article.IsHot,
		ReadCount: article.ReadCount,
		Cid:       article.Cid,
		Category:  BuildCategory(&article.Category),
		CreatedAt: article.CreatedAt.Unix(),
		UpdatedAt: article.UpdatedAt.Unix(),
	}
}

// BuildArticleList 序列化文章列表
func BuildArticleList(items []models.Article) (articleList []ArticleSerializer) {
	for _, item := range items {
		article := BuildArticle(&item)
		articleList = append(articleList, article)
	}
	return articleList
}
