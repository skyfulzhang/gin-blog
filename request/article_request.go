package request

/*
request 实际上就是vo
*/

type ArticleId struct {
	Id int `uri:"id" form:"id" json:"id"  binding:"required,gt=0"`
}

type ArticleCreate struct {
	Title   string `json:"title" form:"title" binding:"required,gt=2,lt=64"`
	Author  string `json:"author" form:"author" binding:"required,gt=2,lt=32"`
	Cover   string `json:"cover" form:"cover" binding:"required,gt=2,lt=126"`
	Content string `json:"content" form:"content" binding:"required,gt=2,lt=256"`
	Cid     int    `json:"cid" form:"cid" binding:"required,gt=0"`
}

type ArticleUpdate struct {
	Title   string `json:"title" form:"title" binding:"required,gt=2,lt=64"`
	Author  string `json:"author" form:"author" binding:"required,gt=2,lt=32"`
	Cover   string `json:"cover" form:"cover" binding:"required,gt=2,lt=126"`
	Content string `json:"content" form:"content" binding:"required,gt=2,lt=256"`
	Cid     int    `json:"cid" form:"cid" binding:"required,gt=0"`
}

type ArticlePagination struct {
	Page int `uri:"page" form:"page" json:"page"`
	Size int `uri:"size" form:"size" json:"size"`
}

type ArticleSearch struct {
	ArticlePagination
	Title string `json:"title" form:"title" `
}
