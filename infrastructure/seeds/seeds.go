package seeds

import (
	"github.com/jinzhu/gorm"
	"tableTennis/domain"
)

func Run() {
	db, err := gorm.Open("mysql", "root:password@/table_tennis?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&domain.Style{})
	db.Create(&domain.Style{Name: "前陣速攻型"})
	db.Create(&domain.Style{Name: "カットマン"})
	db.Create(&domain.Style{Name: "両ハンドドライブ型"})
	db.Create(&domain.Style{Name: "ペン表"})
	db.Create(&domain.Style{Name: "中ペン両ハンドドライブ型"})
	db.Create(&domain.Style{Name: "日ペンドライブ型"})
}
