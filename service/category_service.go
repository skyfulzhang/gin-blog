package service

import (
	"gin-blog/repository"
	"gin-blog/response"
	"gin-blog/serializer"
)

type ICategoryService interface {
	GetCategoryByIdService(id int) response.Response
	GetCategoryListService(page, size int) response.Response
	SearchCategoryService(name string, page, size int) response.Response
	CreateCategoryService(name string) response.Response
	UpdateCategoryService(id int, maps interface{}) response.Response
	ActiveCategoryService(id int) response.Response
	DisableCategoryService(id int) response.Response
	DeleteCategoryService(id int) response.Response
}

type CategoryService struct {
	Repository repository.ICategoryRepository
}

func NewCategoryService() ICategoryService {
	return CategoryService{Repository: repository.NewCategoryRepository()}
}


func (service CategoryService) GetCategoryByIdService(id int) response.Response {
	cate, err := service.Repository.RetrieveCategoryRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_CATEGORY_NOT_EXIST)
	}
	res := serializer.BuildCategory(cate)
	return response.ResponseSuccess(res)
}

func (service CategoryService) GetCategoryListService(page, size int) response.Response {
	cateList, err := service.Repository.ListCategoryRepository(page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildCategoryList(cateList)
	dataMap["total"] = service.Repository.TotalCategoryRepository(nil)
	return response.ResponseSuccess(dataMap)
}

func (service CategoryService) SearchCategoryService(name string, page, size int) response.Response {
	cateList, err := service.Repository.SearchCategoryRepository(name, page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildCategoryList(cateList)
	dataMap["total"] = service.Repository.TotalCategoryRepository(map[string]interface{}{"name": name})
	return response.ResponseSuccess(dataMap)
}

func (service CategoryService) CreateCategoryService(name string) response.Response {
	ok, err := service.Repository.CheckCategoryRepository(name)
	if ok != true {
		return response.ResponseError(response.ERROR_CATEGORY_EXIST)
	}
	cate, err := service.Repository.CreateCategoryRepository(name)
	if err != nil {
		return response.ResponseError(response.ERROR_CREATE_FAILED)
	}
	res := serializer.BuildCategory(cate)
	return response.ResponseSuccess(res)
}

func (service CategoryService) UpdateCategoryService(id int, maps interface{}) response.Response {
	cate, err := service.Repository.UpdateCategoryRepository(id, maps)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := serializer.BuildCategory(cate)
	return response.ResponseSuccess(res)
}

func (service CategoryService) ActiveCategoryService(id int) response.Response {
	err := service.Repository.ActiveCategoryRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetCategoryByIdService(id)
	return res
}

func (service CategoryService) DisableCategoryService(id int) response.Response {
	err := service.Repository.DisableCategoryRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetCategoryByIdService(id)
	return res
}

func (service CategoryService) DeleteCategoryService(id int) response.Response {
	err := service.Repository.DestroyCategoryRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_DELETE_FAILED)
	}
	res := serializer.BuildCategory(nil)
	return response.ResponseSuccess(res)
}
