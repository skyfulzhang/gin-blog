package request

type UserId struct {
	Id int `uri:"id" form:"id" json:"id"  binding:"required,gt=0"`
}

type UserCreate struct {
	Username string `json:"username" form:"username" binding:"required,gt=2,lt=64"`
	Password string `json:"password" form:"password" binding:"required,gt=2,lt=64"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required,gt=2,lt=64"`
	Phone    string `json:"phone" form:"phone" binding:"required,gt=2,lt=64"`
	Email    string `json:"email" form:"email" binding:"required,gt=2,lt=64"`
}

type UserUpdate struct {
	Username string `json:"username" form:"username" binding:"required,gt=2,lt=64"`
	Password string `json:"password" form:"password" binding:"required,gt=2,lt=64"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required,gt=2,lt=64"`
	Phone    string `json:"phone" form:"phone" binding:"required,gt=2,lt=64"`
	Email    string `json:"email" form:"email" binding:"required,gt=2,lt=64"`
}

type UserPagination struct {
	Page int `uri:"page" form:"page" json:"page"`
	Size int `uri:"size" form:"size" json:"size"`
}

type UserSearch struct {
	UserPagination
	Username string `uri:"username" form:"username" json:"username"`
}
