package persistence

import (
	"tableTennis/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence
}

// ユーザ登録 この辺りは gorm 使っているから実装変わりそう
func (up userPersistence) Insert(DB *sql.DB, userID, name, email string) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, name, email)
	return err
}
