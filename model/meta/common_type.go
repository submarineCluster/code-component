// Package meta ...
package meta

import (
	"database/sql/driver"
	"go/types"
	"strings"
)

// StringList TODO
type StringList []string

// Value TODO
func (s StringList) Value() (driver.Value, error) {
	return []byte(strings.Join(s, ",")), nil
}

// Scan TODO
func (s *StringList) Scan(src interface{}) error {
	var body []byte
	switch v := src.(type) {
	case string:
		if len(v) == 0 {
			return nil
		}
		body = []byte(v)
	case []byte:
		if len(v) == 0 {
			return nil
		}
		body = v
	case types.Nil:
		return nil
	default:
		return nil
	}
	*s = strings.Split(string(body), ",")
	return nil
}
