package meta

import (
	"database/sql/driver"
	"encoding/json"
	"go/types"

	"github.com/pkg/errors"
)

//Namespace ...
type Namespace map[NamespaceKey][]string

//Scan ...
func (n *Namespace) Scan(src interface{}) error {
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

	result := Namespace{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return errors.Wrapf(err, "json Unmarshal fail")
	}
	*n = result
	return nil
}

//Value ...
func (n Namespace) Value() (driver.Value, error) {
	if len(n) == 0 {
		return []byte(""), nil
	}
	return json.Marshal(n)
}

//KV ...
type KV map[string]string

//Scan ...
func (kv *KV) Scan(src interface{}) error {
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

	result := KV{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return errors.Wrapf(err, "json Unmarshal fail")
	}
	*kv = result
	return nil
}

//Value ...
func (kv KV) Value() (driver.Value, error) {
	if len(kv) == 0 {
		return []byte(""), nil
	}
	return json.Marshal(kv)
}
