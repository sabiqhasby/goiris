package handler

import (
	"goiris/models"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func ShowBooks(db *gorm.DB) iris.Handler {
	return func(ctx iris.Context) {
		var book []models.Book
		db.Find(&book)

		ctx.JSON(book)
	}

}
