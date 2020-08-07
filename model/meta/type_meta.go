package meta

import "strconv"

//ID id of object
type ID int64

//IsZero ...
func (id ID) IsZero() bool {
	return id == 0
}

// String ...
func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

//Name name of object
type Name string

// String ...
func (n Name) String() string {
	return string(n)
}

//IsZero ...
func (n Name) IsZero() bool {
	return len(n.String()) == 0
}
