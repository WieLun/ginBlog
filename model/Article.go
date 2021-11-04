package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey: CategoryID"`
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"category_id"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Img        string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下所有文章
func GetCateArt(c *gin.Context)  {

}

// 查询单个文章信息
func getArtInfo(c *gin.Context) Article,int  {

}

// 查询文章列表
func GetArt(pageSize, pageNum int) ([]Article,int) {
	var art []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,errmsg.ERROR
	}
	return art,errmsg.SUCCESS
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["category_id"] = data.CategoryID
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	fmt.Println(maps)
	err = db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteArt(id int) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
