package repository

import (
	"errors"
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/utils"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type IUserRepository interface {
	CheckUserRepository(username string) (bool, error)
	CheckLoginRepository(username, password string) (*models.User, error)
	CreateUserRepository(username, password, avatar, phone, email string) (*models.User, error)
	RetrieveUserRepository(id int) (*models.User, error)
	ListUserRepository(pageNum, pageSize int) ([]models.User, error)
	SearchUserRepository(name string, pageNum, pageSize int) ([]models.User, error)
	ActiveUserRepository(id int) error
	DisableUserRepository(id int) error
	ModifyAvatarRepository(id int, avatar string) error
	ModifyPhoneRepository(id int, phone string) error
	ModifyEmailRepository(id int, email string) error
	ResetPwdRepository(id int) error
	TotalUserRepository(maps interface{}) int
	DestroyUserRepository(id int) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() IUserRepository {
	db := dao.GetDB()
	db.AutoMigrate(models.User{})
	return UserRepository{DB: db}
}

// CheckUserRepository 判断用户是否存在
func (rep UserRepository) CheckUserRepository(username string) (bool, error) {
	var user models.User
	err := rep.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if err != nil {
		return false, err
	}
	if user.Id > 0 {
		return false, errors.New("用户已存在")
	}
	return true, nil
}

// CheckLoginRepository  用户登录
func (rep UserRepository) CheckLoginRepository(username, password string) (*models.User, error) {
	var user models.User
	err := rep.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error
	if user.Id == 0 {
		return nil, err
	}
	if utils.EncryMd5(password) != user.Password {
		return nil, errors.New("账号或密码错误")
	}
	if user.State == 0 {
		return nil, errors.New("用户未激活")
	}
	return &user, nil
}

// CreateUserRepository 创建用户
func (rep UserRepository) CreateUserRepository(username, password, avatar, phone, email string) (*models.User, error) {
	user := models.User{Username: username, Password: password, Avatar: avatar, Phone: phone, Email: email}
	err := rep.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "CreateUserRepository"))
		return nil, err
	}
	return &user, nil
}

// RetrieveUserRepository 获取用户详情
func (rep UserRepository) RetrieveUserRepository(id int) (*models.User, error) {
	var user models.User
	err := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "RetrieveUserRepository"))
		return nil, err
	}
	return &user, nil
}

// ListUserRepository 查询全部用户
func (rep UserRepository) ListUserRepository(pageNum, pageSize int) ([]models.User, error) {
	var userList []models.User
	err := rep.DB.Model(&models.User{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "ListUserRepository"))
		return nil, err
	}
	return userList, nil
}

// SearchUserRepository 通过用户名搜索用户
func (rep UserRepository) SearchUserRepository(username string, pageNum, pageSize int) ([]models.User, error) {
	var userList []models.User
	err := rep.DB.Model(&models.User{}).Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userList).Error
	if err != nil {
		zap.L().Error(err.Error(), zap.String("repository", "SearchUserRepository"))
		return nil, err
	}
	return userList, nil
}

// ActiveUserRepository 激活用户
func (rep UserRepository) ActiveUserRepository(id int) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ActiveUserRepository"))
		return err1
	}
	err2 := rep.DB.Model(&user).Update("state", 1).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ActiveUserRepository"))
		return err2
	}
	return nil
}

// DisableUserRepository 禁用用户
func (rep UserRepository) DisableUserRepository(id int) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "DisableUserRepository"))
		return err1
	}
	err2 := rep.DB.Model(&user).Update("state", 2).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "DisableUserRepository"))
		return err2
	}
	return nil
}

// ModifyAvatarRepository 修改手机
func (rep UserRepository) ModifyAvatarRepository(id int, avatar string) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ModifyAvatarRepository"))
		return err1
	}
	err2 := rep.DB.Model(&user).Update("avatar", avatar).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ModifyAvatarRepository"))
		return err2
	}
	return nil
}

// ModifyPhoneRepository 修改手机
func (rep UserRepository) ModifyPhoneRepository(id int, phone string) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ModifyPhoneRepository"))
		return err1
	}
	err2 := rep.DB.Model(&user).Update("phone", phone).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ModifyPhoneRepository"))
		return err2
	}
	return nil
}

// ModifyEmailRepository 修改邮箱
func (rep UserRepository) ModifyEmailRepository(id int, email string) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ModifyEmailRepository"))
		return err1
	}
	err2 := rep.DB.Model(&user).Update("email", email).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ModifyEmailRepository"))
		return err2
	}
	return nil
}

// ResetPwdRepository 重置密码
func (rep UserRepository) ResetPwdRepository(id int) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "ResetPwdRepository"))
		return err1
	}
	hashPwd := utils.EncryMd5("123456")
	err2 := rep.DB.Model(&user).Update("email", hashPwd).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "ResetPwdRepository"))
		return err2
	}
	return err2
}

// TotalUserRepository 条件查询用户数量
func (rep UserRepository) TotalUserRepository(maps interface{}) int {
	var count int
	rep.DB.Model(&models.User{}).Where(maps).Count(&count)
	return count
}

// DestroyUserRepository 删除用户
func (rep UserRepository) DestroyUserRepository(id int) error {
	var user models.User
	err1 := rep.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	if err1 != nil {
		zap.L().Error(err1.Error(), zap.String("repository", "DestroyUserRepository"))
		return err1
	}
	err2 := rep.DB.Model(&models.User{}).Delete(&user).Error
	if err2 != nil {
		zap.L().Error(err2.Error(), zap.String("repository", "DestroyUserRepository"))
		return err2
	}
	return nil
}
