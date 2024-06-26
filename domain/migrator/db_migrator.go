package migrator

import (
	"fmt"

	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
)

func MigrateDb(dbContext core.DbContext) {
	db := dbContext.GetDb()
	if db == nil {
		return
	}
	fmt.Println("migrate db start......")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.OAuthSession{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.BookingSource{})
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.CaseTask{})
	db.AutoMigrate(&model.CaseTaskDetail{})
	fmt.Println("migrate db end.....")
}
