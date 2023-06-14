package model

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

var (
	genders    = make(map[string]struct{})
	status     = make(map[string]struct{})
	attributes = make(map[string]struct{})
)

func init() {
	genders["man"] = struct{}{}
	genders["woman"] = struct{}{}

	status["active"] = struct{}{}
	status["banned"] = struct{}{}
	status["deleted"] = struct{}{}

	attributes["id"] = struct{}{}
	attributes["name"] = struct{}{}
	attributes["surname"] = struct{}{}
	attributes["gender"] = struct{}{}
	attributes["status"] = struct{}{}
	attributes["birth_date"] = struct{}{}
	attributes["creat_date"] = struct{}{}
}

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Gender    string    `json:"gender"`
	Status    string    `json:"status"`
	BirthDate time.Time `json:"birth_date"`
	CreatDate time.Time `json:"creat_date"`
}

func (u *User) Validation() error {
	return nil
}

// FilterUser структура для филтрации
type FilterUser struct {
	// Полное имя (имя + фамилия или фамилия + имя)
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
	f.Desc = form["desc"]                        // ?desc=name,surname
	f.Asc = form["asc"]                          // ?desc=create_date
	number, _ := strconv.Atoi(form.Get("limit")) // ?limit={ number }
	f.Limit = uint(number)
	number, _ = strconv.Atoi(form.Get("offset")) // ?offset={ number }
	f.Offset = uint(number)
}

func (f FilterUser) Validate() error {
	_, ok := genders[f.Gender]
	if !ok {
		return NewValidationError("not correctly gender")
	}
	_, ok = status[f.Status]
	if !ok {
		return NewValidationError("not correctly status")
	}

	// TODO валидация полного имени

	if !validatioAttributes(f.Asc, attributes) {
		return NewValidationError("not correctly asc")
	}

	if !validatioAttributes(f.Desc, attributes) {
		return NewValidationError("not correctly desc")
	}

	return nil
}

func validatioAttributes(att []string, has map[string]struct{}) bool {
	for _, a := range att {
		_, ok := has[a]
		if !ok {
			return false
		}
	}
	return true
}

type UserRepo interface {
	Creat(ctx context.Context, user User) error
	// GetAll(ctx context.Context) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	GetByFilter(ctx context.Context, f FilterUser) ([]User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}

type UserService interface {
	Creat(ctx context.Context, user User) error
	GetByID(ctx context.Context, id uint) (User, error)
	GetByFilter(ctx context.Context, f FilterUser) ([]User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint) error
}
