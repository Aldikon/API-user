package model

import (
	"context"
	"fmt"
	"strings"
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

var (
	ErrUserID        = NewValidationError("not correctly iD")
	ErrUserName      = NewValidationError("not correctly name")
	ErrUserSurname   = NewValidationError("not correctly surname")
	ErrUserGender    = NewValidationError("not correctly gender")
	ErrUserStatus    = NewValidationError("not correctly status")
	ErrUserBirthDate = NewValidationError("not correctly birthDate")
	ErrUserCreatDate = NewValidationError("not correctly creatDate")
	ErrUserFullName  = NewValidationError("not correctly full name")
	ErrUserAsc       = NewValidationError("not correctly asc")
	ErrUserDesc      = NewValidationError("not correctly desc")
)

type User struct {
	ID        uint     `json:"id" redis:"id"`
	Name      string   `json:"name" redis:"name"`
	Surname   string   `json:"surname" redis:"surname"`
	Gender    string   `json:"gender" redis:"gender"`
	Status    string   `json:"status" redis:"status"`
	BirthDate TimeDate `json:"birth_date" redis:"birth_date"`
	CreatDate TimeDate `json:"creat_date" redis:"creat_date"`
}

func SchemaUser(id uint) string {
	return fmt.Sprintf("user:%d", id)
}

func validatioAttributes(att []string) bool {
	for _, a := range att {
		if validationAttributes(a) {
			return false
		}
	}
	return true
}

func validationName(w string) bool {
	return strings.TrimSpace(w) != ""
}

func validationSurname(w string) bool {
	return strings.TrimSpace(w) != ""
}

func validationFullName(w string) bool {
	return strings.TrimSpace(w) != ""
}

func validationGender(w string) bool {
	return has(genders, w) && w != ""
}

func validationStatus(w string) bool {
	return has(status, w) && w != ""
}

func validationAttributes(w string) bool {
	return has(attributes, w) && w != ""
}

func has(h map[string]struct{}, w string) bool {
	_, ok := h[w]
	return ok
}

type UserRepo interface {
	Creat(ctx context.Context, user CreateUser) error
	// GetAll(ctx context.Context) (User, error)
	GetByID(ctx context.Context, id uint) (User, error)
	ListWithFilter(ctx context.Context, f FilterUser) ([]User, error)
	Update(ctx context.Context, user UserUpdate) error
	Delete(ctx context.Context, id uint) error
}

type UserService interface {
	Creat(ctx context.Context, user CreateUser) error
	GetByID(ctx context.Context, id uint) (User, error)
	ListWithFilter(ctx context.Context, f FilterUser) ([]User, error)
	Update(ctx context.Context, user UserUpdate) error
	Delete(ctx context.Context, id uint) error
}
