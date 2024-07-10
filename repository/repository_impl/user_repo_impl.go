package repositoryImpl

import (
	"context"
	"errors"
	"log"
	"trello-api/banana"
	"trello-api/database"
	"trello-api/models"
	"trello-api/repository"

	"github.com/go-sql-driver/mysql"
)

type UserRepositoryImpl struct {
	sql *database.SQL
}

func NewUserRepository(sql *database.SQL) repository.UserRepository {
	return &UserRepositoryImpl{
		sql: sql,
	}
}
func (repo UserRepositoryImpl) SaveUser(context context.Context, user models.User) (models.User, error) {
	statement := `
		INSERT INTO users(user_id, full_name, email, password)
		VALUES (:user_id, :full_name, :email, :password)
	`
	if _, err := repo.sql.Db.NamedExecContext(context, statement, user); err != nil {
		if err, _ := err.(*mysql.MySQLError); err.Number == 1062 {
			log.Println(err.Message)
			return user, banana.ErrUserConflict
		} else {
			log.Println(err.Message)
			return user, errors.New(err.Message)
		}
	}
	return user, nil
}
