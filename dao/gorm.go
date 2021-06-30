package dao

import "gorm.io/gorm"

//Dao ...
type Dao struct {
	*gorm.DB
}

//NewDao ...
func NewDao(db *gorm.DB) *Dao {
	return &Dao{db}
}

//Opt ...
type Opt func(db *gorm.DB) *gorm.DB

//LimitOption ...
func LimitOption(limit int64) Opt {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}
		return db.Limit(int(limit))
	}
}

//OffsetOption ...
func OffsetOption(offset int64) Opt {
	return func(db *gorm.DB) *gorm.DB {
		if offset < 0 {
			return db
		}
		return db.Offset(int(offset))
	}
}
