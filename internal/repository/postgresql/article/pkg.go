package article

import (
	"database/sql/driver"
	"github.com/pkg/errors"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""

		return nil
	}

	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}

	*s = NullString(strVal)

	return nil
}

func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}

	return string(s), nil
}
