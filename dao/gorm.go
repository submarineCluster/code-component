package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

// WhereQueryOption where
func WhereQueryOption(k interface{}, value interface{}) Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(k, value)
	}
}

// BinaryWhereQueryOption  binary where
func BinaryWhereQueryOption(k interface{}, value interface{}) Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("binary %v", k), value)
	}
}

// SelectQueryOption select
func SelectQueryOption(key string) Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(key)
	}
}

// ClausesQueryOption 行锁
func ClausesQueryOption() Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
}

// OrderQueryOption 排序
func OrderQueryOption(values interface{}) Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(values)
	}
}

// RawQueryOption 原生 sql
func RawQueryOption(sql string, values ...interface{}) Opt {
	return func(db *gorm.DB) *gorm.DB {
		return db.Raw(sql, values...)
	}
}
