package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// TimeDate при десериализации с json использует формат времни "2006-01-02"
type TimeDate struct {
	time.Time
}

func (ct *TimeDate) UnmarshalJSON(data []byte) error {
	prefix := "\""
	text := strings.TrimSuffix(string(data), prefix)
	text = strings.TrimPrefix(text, prefix)
	t, err := time.Parse(time.DateOnly, text)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct TimeDate) MarshalJSON() ([]byte, error) {
	if ct.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ct.Time.Format(time.DateOnly))
}

func (ct *TimeDate) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("invalid type %T for CustomDate", value)
	}
	ct.Time = t
	return nil
}

func (ct TimeDate) Value() (driver.Value, error) {
	return ct.Time.Format("2006-01-02"), nil
}
