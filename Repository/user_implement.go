package repository

import (
	"context"
	"database/sql"
	"errors"
	entity "golang-clean-architecture/Entity"
	"strconv"
)

type UserImplement struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) UserInterface {
	return &UserImplement{DB: db}
}

func (repo *UserImplement) Store(ctx context.Context, user entity.User) (entity.User, error) {
	script := "Insert into user(nama,alamat,age,married,date) values(?,?,?,?,?)"
	result, err := repo.DB.ExecContext(ctx, script, user.Name, user.Alamat, user.Age, user.Married, user.Date)

	if err != nil {
		return user, err
	}

	id, _ := result.LastInsertId()
	user.Id = id

	return user, nil
}

func (repo *UserImplement) Delete(ctx context.Context, id int64) error {
	resultId, _ := repo.FindId(ctx, id)

	user := entity.User{}

	if resultId == user {
		return errors.New("Id Tidak Ditemukan")
	}
	script := "delete from user where id=?"
	_, err := repo.DB.ExecContext(ctx, script, id)

	if err != nil {
		return err
	}
	var message error
	message = errors.New("Berhasil Hapus Data")
	return message

}
func (repo *UserImplement) FindId(ctx context.Context, id int64) (entity.User, error) {
	script := "select nama,alamat from user where id=? limit 1"

	result, _ := repo.DB.QueryContext(ctx, script, id)
	user := entity.User{}
	defer result.Close()

	if result.Next() {
		result.Scan(&user.Name, &user.Alamat)
		return user, nil
	} else {
		return user, errors.New("Id: " + strconv.Itoa(int(id)) + " tidak ditemukan")
	}

}
