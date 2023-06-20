package model

type UserUpdate struct {
	ID        uint     `json:"-"`
	Name      string   `json:"name"`
	Surname   string   `json:"surname"`
	Gender    string   `json:"gender"`
	Status    string   `json:"status"`
	BirthDate TimeDate `json:"birth_date"`
}

func (u *UserUpdate) Validation() error {
	// TODO дата рождение не должна быть больше сегоднешнего времени
	return nil
}
