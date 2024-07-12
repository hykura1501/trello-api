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
		INSERT INTO users(user_id, full_name, email, password, avatar, birthday, phone, address, created_at, updated_at)
		VALUES (:user_id, :full_name, :email, :password, :avatar, :birthday, :phone, :address, :created_at, :updated_at)
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

func (repo UserRepositoryImpl) CheckUser(context context.Context, reqLogin models.ReqLogin) (models.User, error) {
	var user models.User
	statement := `
		SELECT * FROM users WHERE email=?
	`
	if err := repo.sql.Db.GetContext(context, &user, statement, reqLogin.Email); err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

func (repo UserRepositoryImpl) GetUser(context context.Context, userId string) (models.User, error) {
	var user models.User
	statement := `
		SELECT * FROM users WHERE user_id=?
	`
	if err := repo.sql.Db.GetContext(context, &user, statement, userId); err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}
