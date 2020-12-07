package dao

import (
	"database/sql"

	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct{
	DB *sql.DB
}

var dao *Dao


type User struct {}

var ErrUserNotFound = errors.New("user not found")

func (d *Dao) GetUserById(userID int) (*User, error) {
	u := &User{}
	if err := dao.DB.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&u); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return u, errors.Wrapf(err, "find user error %v", err)
	}
	return u, nil
}
