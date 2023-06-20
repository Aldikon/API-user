package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/Aldikon/API-user/model"
	"github.com/jackc/pgx/v5"
)

type userDB struct {
	db *pgx.Conn
}

func NewUser(db *pgx.Conn) *userDB {
	return &userDB{
		db: db,
	}
}

func (u *userDB) Creat(ctx context.Context, user model.CreateUser) error {
	query := `
	INSERT INTO Users (name, surname, gender, status, birth_date)
	VALUES ($1, $2, $3, $4, $5);`

	_, err := u.db.Exec(ctx, query,
		user.Name,
		user.Surname,
		user.Gender,
		user.Status,
		user.BirthDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDB) GetByID(ctx context.Context, id uint) (model.User, error) {
	user := model.User{}

	query := `
	SELECT u.id, u.name, u.surname, u.gender, u.status, u.birth_date, u.creat_data 
	FROM Users as u
	WHERE u.id = $1;`

	row := u.db.QueryRow(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Gender,
		&user.Status,
		&user.BirthDate,
		&user.CreatDate,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return user, err
	}

	return user, nil
}

func (u *userDB) ListWithFilter(ctx context.Context, f model.FilterUser) ([]model.User, error) {
	users := make([]model.User, 0)

	query := newFilter(`
	SELECT u.id, u.name, u.surname, u.gender, u.status, u.birth_date, u.creat_data
	FROM users AS u 
	`).withWhere(
		where{
			"gender":                     f.Gender,
			"status":                     f.Status,
			"CONCAT(name, ' ', surname)": f.FullName,
		}).
		withOrderBy(orderBy{
			asc:  f.Asc,
			desc: f.Desc,
		}).
		withLimit(f.Limit).
		withOffset(f.Offset).
		end()

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Gender,
			&user.Status,
			&user.BirthDate,
			&user.CreatDate,
		)
		if err != nil {
			return nil, fmt.Errorf("scan err: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return users, nil
}

func (u *userDB) Update(ctx context.Context, user model.UserUpdate) error {
	query := `
	UPDATE Users
	SET 
		name= $2,
		surname= $3,
		gender= $4,
		status= $5,
		birth_date = $6 
	WHERE id = $1;`

	_, err := u.db.Exec(ctx, query,
		user.ID,
		user.Name,
		user.Surname,
		user.Gender,
		user.Status,
		user.BirthDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDB) Delete(ctx context.Context, id uint) error {
	query := `
	DELETE FROM users as u
	WHERE u.id = $1;`

	_, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
