package model

import (
	"net/url"
	"strconv"
)

type FilterUser struct {
	FullName string
	Gender   string
	Status   string
	Desc     []string
	Asc      []string
	Limit    uint
	Offset   uint
}

func (f *FilterUser) Parse(form url.Values) {
	f.Gender = form.Get("gender")
	f.Status = form.Get("status")
	f.FullName = form.Get("full_name")
	f.Desc = form["desc"]
	f.Asc = form["asc"]
	number, _ := strconv.Atoi(form.Get("limit"))
	f.Limit = uint(number)

	number, _ = strconv.Atoi(form.Get("offset"))
	f.Offset = uint(number)
}

func (f FilterUser) Validate() error {
	if validationGender(f.Gender) {
		return ErrUserGender
	}

	if validationStatus(f.Status) {
		return ErrUserStatus
	}

	if validationFullName(f.FullName) {
		return ErrUserFullName
	}

	if len(f.Asc) != 0 && validatioAttributes(f.Asc) {
		return ErrUserAsc
	}

	if len(f.Asc) != 0 && validatioAttributes(f.Desc) {
		return ErrUserDesc
	}

	return nil
}
