package serializer

import "gin-blog/models"

// CategorySerializer 分类序列化器
type CategorySerializer struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	IsValid   int    `json:"is_valid"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// BuildCategory 序列化分类
func BuildCategory(cate *models.Category) CategorySerializer {
	return CategorySerializer{
		Id:        cate.Id,
		Name:      cate.Name,
		IsValid:   cate.IsValid,
		CreatedAt: cate.CreatedAt.Unix(),
		UpdatedAt: cate.UpdatedAt.Unix(),
	}
}

// BuildCategoryList 序列化文章列表
func BuildCategoryList(items []models.Category) (cateList []CategorySerializer) {
	for _, item := range items {
		cate := BuildCategory(&item)
		cateList = append(cateList, cate)
	}
	return cateList
}
