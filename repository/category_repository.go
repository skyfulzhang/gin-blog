package repository

import (
	"errors"
	"gin-blog/dao"
	"gin-blog/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type ICategoryRepository interface {
	CheckCategoryRepository(name string) (bool, error)
	CreateCategoryRepository(name string) (*models.Category, error)
	RetrieveCategoryRepository(id int) (*models.Category, error)
	ListCategoryRepository(pageNum, pageSize int) ([]models.Category, error)
	UpdateCategoryRepository(id int, maps interface{}) (*models.Category, error)
	SearchCategoryRepository(name string, pageNum, pageSize int) ([]models.Category, error)
	TotalCategoryRepository(maps interface{}) int
	ActiveCategoryRepository(id int) error
	DisableCategoryRepository(id int) error
	DestroyCategoryRepository(id int) error
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() ICategoryRepository {
	db := dao.GetDB()
	db.AutoMigrate(models.Category{})
	return CategoryRepository{DB: db}
}

// CheckCategory 判断分类是否存在
func (rep CategoryRepository) CheckCategoryRepository(name string) (bool, error) {
	var cate models.Category
	err := rep.DB.Model(&models.Category{}).Where("is_valid =2 and name = ?", name).First(&cate).Error
	if err != nil {
		return false, err
	}
	if cate.Id > 0 {
		return false, errors.New("分类已存在")
	}
	return true, nil
}

// CreateCategory  创建分类
func (rep CategoryRepository) CreateCategoryRepository(name string) (*models.Category, error) {
	cate := models.Category{Name: name}
	err := rep.DB.Model(&models.Category{}).Create(&cate).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "CreateCategoryRepository"))
		return nil, err
	}
	return &cate, nil
}

// RetrieveCategoryRepository 查询分类
func (rep CategoryRepository) RetrieveCategoryRepository(id int) (*models.Category, error) {
	var cate models.Category
	err := rep.DB.Model(&models.Category{}).Where("is_valid =2 and id = ?", id).First(&cate).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "RetrieveCategory"))
		return nil, err
	}
	return &cate, nil
}

// ListCategoryRepository 查询全部分类
func (rep CategoryRepository) ListCategoryRepository(pageNum, pageSize int) ([]models.Category, error) {
	var cateList []models.Category
	err := rep.DB.Model(&models.Category{}).Where("is_valid =2").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateList).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "ListCategory"))
		return nil, err
	}
	return cateList, nil
}

// UpdateCategoryRepository 更新分类
func (rep CategoryRepository) UpdateCategoryRepository(id int, maps interface{}) (*models.Category, error) {
	var cate models.Category
	err1 := rep.DB.Model(&models.Category{}).Where("id = ?", id).First(&cate).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "UpdateCategoryRepository"))
		return nil, err1
	}
	err2 := rep.DB.Model(&cate).Updates(maps).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "UpdateCategoryRepository"))
		return nil, err2
	}
	return &cate, nil
}

// SearchCategoryRepository 通过分类名称搜索分类
func (rep CategoryRepository) SearchCategoryRepository(name string, pageNum, pageSize int) ([]models.Category, error) {
	var cateList []models.Category
	err := rep.DB.Model(&models.Category{}).Where("is_valid =2 and name like ?", "%"+name+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateList).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "SearchCategoryRepository"))
		return nil, err
	}
	return cateList, nil
}

// TotalCategoryRepository 查询全部分类
func (rep CategoryRepository) TotalCategoryRepository(maps interface{}) int {
	var count int
	rep.DB.Model(&models.Category{}).Where(maps).Count(&count)
	return count
}

// ActiveCategoryRepository 激活分类
func (rep CategoryRepository) ActiveCategoryRepository(id int) error {
	var cate models.Category
	err1 := rep.DB.Model(&models.Category{}).Where("id = ?", id).First(&cate).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ActiveCategoryRepository"))
		return err1
	}
	err2 := rep.DB.Model(&cate).Update("is_valid", 1).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ActiveCategoryRepository"))
		return err2
	}
	return nil
}

// DisableCategoryRepository 失效分类
func (rep CategoryRepository) DisableCategoryRepository(id int) error {
	var cate models.Category
	err1 := rep.DB.Model(&models.Category{}).Where("id = ?", id).First(&cate).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "DisableCategoryRepository"))
		return err1
	}
	err2 := rep.DB.Model(&cate).Update("is_valid", 2).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "DisableCategoryRepository"))
		return err2
	}
	return nil
}

// DestroyCategoryRepository  删除分类
func (rep CategoryRepository) DestroyCategoryRepository(id int) error {
	var cate models.Category
	err1 := rep.DB.Model(&models.Category{}).Where("id = ?", id).First(&cate).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "DestroyCategoryRepository"))
		return err1
	}
	err2 := rep.DB.Model(&models.Category{}).Delete(&cate).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "DestroyCategoryRepository"))
		return err2
	}
	return nil
}
