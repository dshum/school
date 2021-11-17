package utils

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullString sql.NullString

func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(x.String)
}

func (x *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*x = NullString{s.String, false}
	} else {
		*x = NullString{s.String, true}
	}

	return nil
}
