package books

import (
	_ "fmt"

	"github.com/jinzhu/gorm"
	"github.com/proyecto/backend_go/common"
	"github.com/proyecto/backend_go/users"
)

type BookModel struct {
	gorm.Model
	Slug        string `gorm:"unique_index"`
	Title       string
	Description string `gorm:"size:2048"`
	Category    string
	Author      string
	Price       uint
	Comments    []CommentBookModel `gorm:"ForeignKey:BookID"`
}

type CommentBookModel struct {
	gorm.Model
	Book     BookModel
	BookID   uint
	Author   BookUserModel
	AuthorID uint
	Body     string `gorm:"size:2048"`
}
type BookUserModel struct {
	gorm.Model
	UserModel      users.UserModel
	UserModelID    uint
	BookModels     []BookModel     `gorm:"ForeignKey:AuthorID"`
	FavoriteModels []FavoriteModel `gorm:"ForeignKey:FavoriteByID"`
}
type FavoriteModel struct {
	gorm.Model
	Favorite     BookModel
	FavoriteID   uint
	FavoriteBy   BookUserModel
	FavoriteByID uint
}

func GetBookUserModel(userModel users.UserModel) BookUserModel {
	var bookUserModel BookUserModel
	if userModel.ID == 0 {
		return bookUserModel
	}
	db := common.GetDB()
	db.Where(&BookUserModel{
		UserModelID: userModel.ID,
	}).FirstOrCreate(&bookUserModel)
	bookUserModel.UserModel = userModel
	return bookUserModel
}
func (self *BookModel) getComments() error {
	db := common.GetDB()
	tx := db.Begin()
	tx.Model(self).Related(&self.Comments, "Comments")
	for i, _ := range self.Comments {
		tx.Model(&self.Comments[i]).Related(&self.Comments[i].Author, "Author")
		tx.Model(&self.Comments[i].Author).Related(&self.Comments[i].Author.UserModel)
	}
	err := tx.Commit().Error
	return err
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

func DeleteCommentBookModel(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(CommentBookModel{}).Error
	return err
}
func (book BookModel) favoritesCount() uint {
	db := common.GetDB()
	var count uint
	db.Model(&FavoriteModel{}).Where(FavoriteModel{
		FavoriteID: book.ID,
	}).Count(&count)
	return count
}

func (book BookModel) isFavoriteBy(user BookUserModel) bool {
	db := common.GetDB()
	var favorite FavoriteModel
	db.Where(FavoriteModel{
		FavoriteID:   book.ID,
		FavoriteByID: user.ID,
	}).First(&favorite)
	return favorite.ID != 0
}

func (book BookModel) favoriteBy(user BookUserModel) error {
	db := common.GetDB()
	var favorite FavoriteModel
	err := db.FirstOrCreate(&favorite, &FavoriteModel{
		FavoriteID:   book.ID,
		FavoriteByID: user.ID,
	}).Error
	return err
}
func (book BookModel) unFavoriteBy(user BookUserModel) error {
	db := common.GetDB()
	err := db.Where(FavoriteModel{
		FavoriteID:   book.ID,
		FavoriteByID: user.ID,
	}).Delete(FavoriteModel{}).Error
	return err
}
