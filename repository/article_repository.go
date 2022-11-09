package repository

import (
	"fmt"
	"gin-blog/dao"
	"gin-blog/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type IArticleRepository interface {
	CreateArticleRepository(title, author, cover, content string, cid int) (*models.Article, error)
	RetrieveArticleRepository(id int) (*models.Article, error)
	GetCateArticleRepository(id, pageNum, pageSize int) ([]models.Article, int, error)
	ListArticleRepository(pageNum, pageSize int) ([]models.Article, int, error)
	UpdateArticleRepository(id int, maps interface{}) (*models.Article, error)
	FilterArticleRepository(title, content string, status, is_hot, cid, page, size int) ([]models.Article, int, error)
	UpdateReadCountRepository(id int) error
	PutOnArticleRepository(id int) error
	PullOffArticleRepository(id int) error
	HotArticleRepository(id int) error
	DestroyArticleRepository(id int) error
}

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository() IArticleRepository {
	db := dao.GetDB()
	db.AutoMigrate(models.Article{})
	return ArticleRepository{DB: db}
}

// CreateArticleRepository 创建文章
func (rep ArticleRepository) CreateArticleRepository(title, author, cover, content string, cid int) (*models.Article, error) {
	article := models.Article{
		Title:   title,
		Author:  author,
		Cover:   cover,
		Content: content,
		Cid:     cid,
	}
	err := rep.DB.Model(&models.Article{}).Preload("Category").Create(&article).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "CreateArticle"))
		return nil, err
	}
	return &article, nil
}

// RetrieveArticleRepository 获取文章详情
func (rep ArticleRepository) RetrieveArticleRepository(id int) (*models.Article, error) {
	var article models.Article
	err := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "RetrieveArticle"))
		return nil, err
	}
	fmt.Println("article:", article)
	return &article, err
}

// GetCateArticleRepository 获取分类下的文章
func (rep ArticleRepository) GetCateArticleRepository(id, pageNum, pageSize int) ([]models.Article, int, error) {
	var count int
	var cateArtList []models.Article
	err := rep.DB.Model(&models.Article{}).Where("cid = ?", id).Preload("Category").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("id desc").Find(&cateArtList).Count(&count).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "GetCateArticle"))
		return nil, count, err
	}
	return cateArtList, count, nil
}

// ListArticleRepository 查询全部文章
func (rep ArticleRepository) ListArticleRepository(pageNum, pageSize int) ([]models.Article, int, error) {
	var count int
	var articleList []models.Article
	err := rep.DB.Model(&models.Article{}).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("id desc").Find(&articleList).Count(&count).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "ListArticleRepository"))
		return nil, count, err
	}
	return articleList, count, nil
}

// ExactArticleRepository 精准查询文章 post请求 暂时不用
func (rep ArticleRepository) ExactArticleRepository(where map[string]interface{}, pageNum, pageSize int) ([]models.Article, int, error) {
	var count int
	var articleList []models.Article
	err := rep.DB.Model(&models.Article{}).Preload("Category").Where(where).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("id desc").Find(&articleList).Count(&count).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "ExactArticleRepository"))
		return nil, count, err
	}
	return articleList, count, nil
}

// UpdateArticleRepository 修改文章
func (rep ArticleRepository) UpdateArticleRepository(id int, maps interface{}) (*models.Article, error) {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "UpdateArticleRepository"))
		return nil, err1
	}
	err2 := rep.DB.Model(&article).Updates(maps).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "UpdateArticleRepository"))
		return nil, err2
	}
	return &article, nil
}

// UpdateReadCountRepository 更新阅读数量
func (rep ArticleRepository) UpdateReadCountRepository(id int) error {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		return err1
	}
	err2 := rep.DB.Model(&article).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1)).Error
	if err2 != nil {
		return err2
	}
	return nil
}

// FilterArticleRepository  通过条件筛选文章
func (rep ArticleRepository) FilterArticleRepository(title, content string, status, is_hot, cid, page, size int) ([]models.Article, int, error) {
	var count int
	var articleList []models.Article
	tx := rep.DB.Model(&models.Article{}).Preload("Category")
	if title != "" {
		tx = tx.Where("title like ?", "%"+title+"%")
	}
	if content != "" {
		tx = tx.Where("content like ?", "%"+content+"%")
	}
	if status > 0 {
		tx = tx.Where("status = ?", status)
	}
	if is_hot > 0 {
		tx = tx.Where("is_hot = ?", is_hot)
	}
	if cid > 0 {
		tx = tx.Where("cid = ?", cid)
	}
	if page > 0 && size > 0 {
		tx = tx.Limit(size).Offset((page - 1) * size)
	}
	err := tx.Order("id desc").Find(&articleList).Count(&count).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "FilterArticleRepository"))
		return nil, count, err
	}
	return articleList, count, nil
}

// PutOnArticleRepository 文章上架
func (rep ArticleRepository) PutOnArticleRepository(id int) error {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "PutOnArticleRepository"))
		return err1
	}
	err2 := rep.DB.Model(&article).Update("status", 2).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "PutOnArticleRepository"))
		return err2
	}
	return nil
}

// PullOffArticleRepository 文章下架
func (rep ArticleRepository) PullOffArticleRepository(id int) error {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "PullOffArticleRepository"))
		return err1
	}
	err2 := rep.DB.Model(&article).Update("status", 3).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "PullOffArticleRepository"))
		return err2
	}
	return nil
}

// HotArticleRepository 文章修改为热门
func (rep ArticleRepository) HotArticleRepository(id int) error {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "HotArticleRepository"))
		return err1
	}
	err2 := rep.DB.Model(&article).Update("is_hot", 1).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "HotArticleRepository"))
		return err2
	}
	return nil
}

// DestroyArticleRepository 删除文章
func (rep ArticleRepository) DestroyArticleRepository(id int) error {
	var article models.Article
	err1 := rep.DB.Model(&models.Article{}).Where("id = ?", id).Preload("Category").First(&article).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "DestroyArticleRepository"))
		return err1
	}
	err2 := rep.DB.Model(&models.Article{}).Delete(&article).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "DestroyArticleRepository"))
		return err2
	}
	return nil
}
