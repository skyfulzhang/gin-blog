package service

import (
	"gin-blog/repository"
	"gin-blog/response"
	"gin-blog/serializer"
)

type IArticleService interface {
	GetArticleByIdService(id int) response.Response
	GetCateArticleService(id, page, size int) response.Response
	GetArticleListService(page, size int) response.Response
	CreateArticleService(title, author, cover, content string, cid int) response.Response
	UpdateArticleService(id int, maps interface{}) response.Response
	PutOnArticleService(id int) response.Response
	PullOffArticleService(id int) response.Response
	HotArticleService(id int) response.Response
	FilterArticleService(title, content string, status, is_hot, cid, page, size int) response.Response
	DeleteArticleService(id int) response.Response
}

type ArticleService struct {
	Repository repository.IArticleRepository
}

func NewArticleService() IArticleService {
	return ArticleService{Repository: repository.NewArticleRepository()}
}

func (service ArticleService) GetArticleByIdService(id int) response.Response {
	article, err1 := service.Repository.RetrieveArticleRepository(id)
	if err1 != nil {
		return response.ResponseError(response.ERROR_ARTICLE_NOT_EXIST)
	}
	err2 := service.Repository.UpdateReadCountRepository(id)
	if err2 != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := serializer.BuildArticle(article)
	return response.ResponseSuccess(res)
}

func (service ArticleService) GetCateArticleService(id, page, size int) response.Response {
	articleList, count, err := service.Repository.GetCateArticleRepository(id, page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_ARTICLE_NOT_EXIST)
	}
	res := serializer.BuildArticleList(articleList)
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildArticleList(articleList)
	dataMap["total"] = count
	return response.ResponseSuccess(res)
}

func (service ArticleService) GetArticleListService(page, size int) response.Response {
	articleList, count, err := service.Repository.ListArticleRepository(page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildArticleList(articleList)
	dataMap["total"] = count
	return response.ResponseSuccess(dataMap)
}

func (service ArticleService) FilterArticleService(title, content string, status, is_hot, cid, page, size int) response.Response {
	articleList, count, err := service.Repository.FilterArticleRepository(title, content, status, is_hot, cid, page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildArticleList(articleList)
	dataMap["total"] = count
	return response.ResponseSuccess(dataMap)
}

func (service ArticleService) CreateArticleService(title, author, cover, content string, cid int) response.Response {
	article, err := service.Repository.CreateArticleRepository(title, author, cover, content, cid)
	if err != nil {
		return response.ResponseError(response.ERROR_CREATE_FAILED)
	}
	res := serializer.BuildArticle(article)
	return response.ResponseSuccess(res)
}

func (service ArticleService) UpdateArticleService(id int, maps interface{}) response.Response {
	article, err := service.Repository.UpdateArticleRepository(id, maps)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := serializer.BuildArticle(article)
	return response.ResponseSuccess(res)
}

func (service ArticleService) PutOnArticleService(id int) response.Response {
	err := service.Repository.PutOnArticleRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetArticleByIdService(id)
	return res
}

func (service ArticleService) PullOffArticleService(id int) response.Response {
	err := service.Repository.PullOffArticleRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetArticleByIdService(id)
	return res
}

func (service ArticleService) HotArticleService(id int) response.Response {
	err := service.Repository.HotArticleRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetArticleByIdService(id)
	return res
}

func (service ArticleService) DeleteArticleService(id int) response.Response {
	err := service.Repository.DestroyArticleRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_DELETE_FAILED)
	}
	res := serializer.BuildArticle(nil)
	return response.ResponseSuccess(res)
}
