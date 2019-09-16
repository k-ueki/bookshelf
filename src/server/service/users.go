package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-ueki/app2/src/server/model"
	"github.com/k-ueki/app2/src/server/repository"
)

type User struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *User {
	return &User{db}
}
func (u *User) Index(uid string) (interface{}, error) {
	res := model.UserResp{}

	usr, err := repository.SelectUserByUid(u.DB, uid)
	if err != nil {
		return nil, err
	}

	books, err := repository.SelectBookByUserId(u.DB, usr.Id)
	if err != nil {
		return nil, err
	}

	coms, err := repository.GetAllByUid(u.DB, uid)
	if err != nil {
		return nil, err
	}

	res.Id = usr.Id
	res.Name = usr.Name
	res.Books = *books
	res.Communities = *coms

	return res, nil
}

func main() {}
