package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(32);not null;comment:'用户名'" json:"username" `
	Password string `gorm:"type:varchar(64);not null;comment:'密码'" json:"-"`
	Avatar   string `gorm:"type:varchar(64);not null;comment:'头像'" json:"avatar"`
	Phone    string `gorm:"type:varchar(64);not null;comment:'手机'" json:"phone"`
	Email    string `gorm:"type:varchar(64);not null;comment:'邮箱'" json:"email"`
	State    int    `gorm:"type:int;not null;comment:'是否激活：1-是 2-否'" json:"state"`
}

// 自定义表名
func (User) TableName() string {
	return "user"
}

// MD5加密处理
func EncryMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// BeforeCreate 密码加密
func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	user.Password = EncryMd5(user.Password)
	return nil
}

func (user *User) SetUsername(username string) {
	user.Username = username
}

func (user *User) SetPassword(password string) {
	user.Password = EncryMd5(password)
}

func (user *User) SetAvatar(avatar string) {
	user.Avatar = avatar
}

func (user *User) SetPhone(phone string) {
	user.Phone = phone
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) SetContent(status int) {
	user.State = status
}

//// 使用事务删除角色
//func Delete(roleids []uint64) error {
//	tx := dao.GetDB().Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//		}
//	}()
//	if err := tx.Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	if err := tx.Where("id in (?)", roleids).Delete(&User{}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	if err := tx.Where("role_id in (?)", roleids).Delete(&User{}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	return tx.Commit().Error
//}
//
//// 事务封装
//func Transaction(funcs ...func(db *gorm.DB) error) (err error) {
//	tx := dao.GetDB().Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			tx.Rollback()
//			err = fmt.Errorf("%v", err)
//		}
//	}()
//	for _, f := range funcs {
//		err = f(tx)
//		if err != nil {
//			tx.Rollback()
//			return
//		}
//	}
//	err = tx.Commit().Error
//	return
//}
