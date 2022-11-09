package service

import (
	"gin-blog/repository"
	"gin-blog/response"
	"gin-blog/serializer"
	"gin-blog/utils"
)

type IUserService interface {
	GetUserByIdService(id int) response.Response
	GetUserListService(page, size int) response.Response
	SearchUserService(username string, page, size int) response.Response
	UserLoginService(username, password string) response.Response
	CreateUserService(username, password, avatar, phone, email string) response.Response
	UpdateUserService(id int, avatar, phone, email string) response.Response
	ActiveUserService(id int) response.Response
	DisableUserService(id int) response.Response
	ResetPwdService(id int) response.Response
	DeleteUserService(id int) response.Response
}

type UserService struct {
	Repository repository.IUserRepository
}

func NewUserService() IUserService {
	return UserService{Repository: repository.NewUserRepository()}
}

func (service UserService) GetUserByIdService(id int) response.Response {
	user, err := service.Repository.RetrieveUserRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_USER_NOT_EXIST)
	}
	res := serializer.BuildUser(user)
	return response.ResponseSuccess(res)
}

func (service UserService) GetUserListService(page, size int) response.Response {
	userList, err := service.Repository.ListUserRepository(page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildUserList(userList)
	dataMap["total"] = service.Repository.TotalUserRepository(nil)
	return response.ResponseSuccess(dataMap)
}

func (service UserService) SearchUserService(username string, page, size int) response.Response {
	userList, err := service.Repository.SearchUserRepository(username, page, size)
	if err != nil {
		return response.ResponseError(response.ERROR_QUERY_FAILED)
	}
	var dataMap = make(map[string]interface{})
	dataMap["data"] = serializer.BuildUserList(userList)
	dataMap["total"] = service.Repository.TotalUserRepository(map[string]interface{}{"username": username})
	return response.ResponseSuccess(dataMap)
}

func (service UserService) UserLoginService(username, password string) response.Response {
	user, err := service.Repository.CheckLoginRepository(username, password)
	if err != nil {
		return response.ResponseError(response.ERROR_USER_NOT_EXIST)
	}
	token, err := utils.GenerateToken(user.Id, user.Username, 101344001)
	if err != nil {
		return response.ResponseError(response.ERROR_TOKEN_GENERATE_FAILED)
	}
	res := response.ResponseSuccess(token)
	return res
}

func (service UserService) CreateUserService(username, password, avatar, phone, email string) response.Response {
	ok, err := service.Repository.CheckUserRepository(username)
	if ok != true {
		return response.ResponseError(response.ERROR_USER_EXIST)
	}
	user, err := service.Repository.CreateUserRepository(username, password, avatar, phone, email)
	if err != nil {
		return response.ResponseError(response.ERROR_CREATE_FAILED)
	}
	res := serializer.BuildUser(user)
	return response.ResponseSuccess(res)
}

func (service UserService) UpdateUserService(id int, avatar, phone, email string) response.Response {
	err1 := service.Repository.ModifyAvatarRepository(id, avatar)
	err2 := service.Repository.ModifyPhoneRepository(id, phone)
	err3 := service.Repository.ModifyEmailRepository(id, email)
	if err1 != nil || err2 != nil || err3 != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}

	res := service.GetUserByIdService(id)
	return res
}

func (service UserService) ActiveUserService(id int) response.Response {
	err := service.Repository.ActiveUserRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetUserByIdService(id)
	return res
}

func (service UserService) DisableUserService(id int) response.Response {
	err := service.Repository.DisableUserRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := service.GetUserByIdService(id)
	return res
}

func (service UserService) ResetPwdService(id int) response.Response {
	err := service.Repository.ResetPwdRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_UPDATE_FAILED)
	}
	res := serializer.BuildUser(nil)
	return response.ResponseSuccess(res)
}

func (service UserService) DeleteUserService(id int) response.Response {
	err := service.Repository.DestroyUserRepository(id)
	if err != nil {
		return response.ResponseError(response.ERROR_DELETE_FAILED)
	}
	res := serializer.BuildUser(nil)
	return response.ResponseSuccess(res)
}
