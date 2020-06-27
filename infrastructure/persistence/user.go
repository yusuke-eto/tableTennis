package persistence

import (
	"tableTennis/domain"
	"tableTennis/domain/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// ユーザ登録 この辺りは gorm 使っているから実装変わりそう
func (up userPersistence) Insert(user domain.User) error {
	db, err := gorm.Open("mysql", "root:password@/table_tennis?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.Create(&user)

	return err
}
