package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	BaseModel
	Title     string   `gorm:"type:varchar(64);not null;comment:'文章标题'" json:"title"`
	Author    string   `gorm:"type:varchar(32);not null;comment:'文章作者'" json:"author"`
	Cover     string   `gorm:"type:varchar(126);not null;comment:'文章封面'" json:"cover"`
	Content   string   `gorm:"type:longtext;not null;comment:'文章内容'" json:"content"`
	Status    int      `gorm:"type:int;not null;default:1;comment:'文章状态：1-待审核 2-已上架 3-已下架'" json:"status"`
	IsHot     int      `gorm:"type:int;not null;default:1;comment:'是否热门：1-否 2-是'" json:"is_hot"`
	ReadCount int      `gorm:"type:int;not null;default:1;comment:'阅读数量'" json:"read_count"`
	Cid       int      `gorm:"type:int;not null;comment:'分类名称'" json:"cid"`
	Category  Category `gorm:"ForeignKey:Cid"`
}

// 自定义表名
func (Article) TableName() string {
	return "article"
}

//BeforeCreate 在创建之前，先把创建时间赋值
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

//BeforeUpdate 在更新之前，先把更新时间赋值
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (article *Article) SetTitle(title string) {
	article.Title = title
}

func (article *Article) SetAuthor(author string) {
	article.Author = author
}

func (article *Article) SetCover(cover string) {
	article.Cover = cover
}

func (article *Article) SetName(title string) {
	article.Title = title
}

func (article *Article) SetContent(content string) {
	article.Content = content
}

func (article *Article) SetStatus(status int) {
	article.Status = status
}

func (article *Article) SetIsHot(isHot int) {
	article.IsHot = isHot
}

func (article *Article) SetReadCount(readCount int) {
	article.ReadCount = readCount
}

func (article *Article) SetCid(cid int) {
	article.Cid = cid
}
