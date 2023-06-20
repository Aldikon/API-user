package model

import "time"

type CreateUser struct {
	Name      string   `json:"name"`
	Surname   string   `json:"surname"`
	Gender    string   `json:"gender"`
	Status    string   `json:"status"`
	BirthDate TimeDate `json:"birth_date"`
}

func (u *CreateUser) Validation() error {
	// TODO дата рождение не должна быть больше сегоднешнего времени

	if !validationName(u.Name) {
		return ErrUserName
	}

	if !validationSurname(u.Surname) {
		return ErrUserStatus
	}

	if !validationGender(u.Gender) {
		return ErrUserGender
	}

	if !validationStatus(u.Status) {
		return ErrUserStatus
	}

	if !u.BirthDate.Before(time.Now()) {
		return ErrUserBirthDate
	}

	return nil
}
