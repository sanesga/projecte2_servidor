package books

import (
	_ "fmt"

	"github.com/backend_go/common"
	"github.com/jinzhu/gorm"
)

type BookModel struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Category    string
	Author      string
	Price       uint
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneBook(condition interface{}) (BookModel, error) {
	db := common.GetDB()
	var model BookModel
	tx := db.Begin() //tx == transaction
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func getAllBooks() ([]BookModel, error) {
	db := common.GetDB()
	var models []BookModel
	err := db.Find(&models).Error
	return models, err
}

func (model *BookModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func DeleteBookModel(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(BookModel{}).Error
	return err
}
