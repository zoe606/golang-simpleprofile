package repositories

import (
	"context"
	"database/sql"
	"errors"
	"simpleProfile/helpers"
	"simpleProfile/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile {
	qry := "INSERT INTO USER(username,password,firstname,lastname,age,phone,address,email) values (?,?,?,?,?,?,?,?)"
	res, err := tx.ExecContext(ctx, qry, profile.Username, profile.Password, profile.Firstname, profile.Lastname, profile.Age, profile.Phone, profile.Address, profile.Email)
	helpers.PanicIfError(err)

	id, err := res.LastInsertId()
	helpers.PanicIfError(err)

	profile.Id = int(id)

	return profile
}

func (u UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, profile domain.Profile) (domain.Profile, error) {
	query := "SELECT id,username,password FROM USER where username = ?"
	rows, err := tx.QueryContext(ctx, query, profile.Username)
	helpers.PanicIfError(err)
	defer rows.Close()

	user := domain.Profile{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		helpers.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (u UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile {
	query := "UPDATE USER set username = ?,password = ?,firstname = ?,lastname = ?,age = ?,phone = ?,address = ?,email = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, profile.Username, profile.Password, profile.Firstname, profile.Lastname, profile.Age, profile.Phone, profile.Address, profile.Email, profile.Id)
	helpers.PanicIfError(err)

	return profile
}

func (u UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, profileId int) (domain.Profile, error) {
	query := "SELECT id,username,password,firstname,lastname,age,phone,address,email FROM USER where id = ?"
	rows, err := tx.QueryContext(ctx, query, profileId)
	helpers.PanicIfError(err)
	defer rows.Close()

	user := domain.Profile{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Age, &user.Phone, &user.Address, &user.Email)
		helpers.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (u UserRepositoryImpl) Logout(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile {
	//TODO implement me
	panic("implement me")
}
