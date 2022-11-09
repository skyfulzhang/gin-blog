package request

type CategoryId struct {
	Id int `uri:"id" form:"id" json:"id"  binding:"required,gt=0"`
}

type CategoryCreate struct {
	Name    string `json:"name" form:"name" binding:"required,gt=2,lt=64"`
	IsValid int    `json:"is_valid" form:"name" binding:"required,oneof=0 1"`
}

type CategoryUpdate struct {
	Name    string `json:"name" form:"name" binding:"required,gt=2,lt=64"`
	IsValid int    `json:"is_valid" form:"name" binding:"required,oneof=0 1"`
}

type CategoryPagination struct {
	Page int `uri:"page" form:"page" json:"page"`
	Size int `uri:"size" form:"size" json:"size"`
}

type CategorySearch struct {
	CategoryPagination
	Name string `uri:"name" form:"name" json:"name"`
}
