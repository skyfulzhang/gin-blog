package serializer

import (
	"gin-blog/models"
)

// UserSerializer 用户序列化器
type UserSerializer struct {
	Id        uint64 `json:"id"`
	UserName  string `json:"user_name"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	State     int    `json:"state"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// BuildUser 序列化用户
func BuildUser(user *models.User) UserSerializer {
	return UserSerializer{
		Id:        user.Id,
		UserName:  user.Username,
		State:     user.State,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}

// BuildUserList 序列化用户列表
func BuildUserList(items []models.User) (userList []UserSerializer) {
	for _, item := range items {
		user := BuildUser(&item)
		userList = append(userList, user)
	}
	return userList
}
