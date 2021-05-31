package option

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//ListOption list操作通用配置项
type ListOption struct {

	//FlagBit 标志位
	FlagBit FlagBit `json:"flagBit,omitempty"`
	// Page 分页相关
	Page Page `json:"page,omitempty"`

	// OrderByList 排序
	OrderByList []OrderBy `json:"sort,omitempty"`

	// Filter 只返回指定的字段，filter 为空则全部返回
	Filter []string `json:"filter,omitempty"`

	// ExtendMap 拓展字段
	ExtendMap map[string]interface{} `json:"extendMap,omitempty"`
}

//NewListOption ...
func NewListOption(opts ...ListOpt) *ListOption {
	o := &ListOption{
		FlagBit: 0,
		Page: Page{
			Limit:  10,
			Offset: 0,
		},
		OrderByList: []OrderBy{
			{
				Field:       "id",
				OrderByType: OrderByTypeDesc,
			},
		},
		Filter:    nil,
		ExtendMap: nil,
	}

	for _, opt := range opts {
		o = opt(o)
	}
	return o
}

//Page ...
type Page struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (o *ListOption) setLimit(limit int64) {
	o.Page.Limit = limit
}

func (o *ListOption) setOffset(offset int64) {
	o.Page.Offset = offset
}

func (o *ListOption) setExtendMapKV(k string, v interface{}) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]interface{})
	}
	o.ExtendMap[k] = v
}

//OrderByType ...
type OrderByType string

//OrderBy ...
type OrderBy struct {
	Field       string      `json:"field"`
	OrderByType OrderByType `json:"orderByType"`
}

// const ...
const (
	OrderByTypeDesc OrderByType = "desc"
	OrderByTypeAsc  OrderByType = "asc"
)

//SetOrderBy ...
func (o *ListOption) SetOrderBy(by OrderBy) {
	o.OrderByList = append(o.OrderByList, by)
}

//ListOpt ...
type ListOpt func(o *ListOption) *ListOption

//WithDeleteListOption ...
func WithDeleteListOption() ListOpt {
	return func(o *ListOption) *ListOption {
		o.FlagBit |= DeleteFlagBit
		return o
	}
}

//LimitListOption ...
func LimitListOption(limit int64) ListOpt {
	return func(o *ListOption) *ListOption {
		o.setLimit(limit)
		return o
	}
}

//OffsetListOption ...
func OffsetListOption(offset int64) ListOpt {
	return func(o *ListOption) *ListOption {
		o.setOffset(offset)
		return o
	}
}

//SetExtendMapListOption ...
func SetExtendMapListOption(key string, value interface{}) ListOpt {
	return func(o *ListOption) *ListOption {
		o.setExtendMapKV(key, value)
		return o
	}
}

//SettingDBModel 配置 dbModel
func SettingDBModel(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {
	if option == nil || dbModel == nil {
		return nil, errors.Errorf("invalid option")
	}

	var err error
	for _, f := range DBModelFilterHandlerFactory {
		dbModel, err = f(dbModel, option)
		if err != nil {
			return nil, errors.Wrapf(err, "filter")
		}
	}
	return dbModel, nil
}

//DBModelFilterHandler ...
type DBModelFilterHandler func(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error)

var (
	//DBModelFilterHandlerFactory ...
	DBModelFilterHandlerFactory = []DBModelFilterHandler{
		DBModelFilterFlagBit,
		DBModelFilterPage,
		DBModelFilterOrderBy,
		DBModelFilterFieldList,
		DBModelFilterExtendMap,
	}
)

//DBModelFilterFlagBit ...
func DBModelFilterFlagBit(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {
	if option.FlagBit&DeleteFlagBit > 0 {
		dbModel = dbModel.Where("delete_timestamp > ?", 0)
	}
	return dbModel, nil
}

//DBModelFilterPage ...
func DBModelFilterPage(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {

	if option.Page.Limit > 0 {
		dbModel = dbModel.Limit(option.Page.Limit)
	}

	if option.Page.Offset >= 0 {
		dbModel = dbModel.Offset(option.Page.Offset)
	}

	return dbModel, nil
}

//DBModelFilterOrderBy ...
func DBModelFilterOrderBy(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {

	for _, ob := range option.OrderByList {
		dbModel = dbModel.Order(fmt.Sprintf("%v %v", ob.Field, ob.OrderByType))
	}

	return dbModel, nil
}

//DBModelFilterFieldList ...
func DBModelFilterFieldList(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {

	if len(option.Filter) > 0 {
		dbModel = dbModel.Select(strings.Join(option.Filter, ","))
	}

	return dbModel, nil
}

//DBModelFilterExtendMap ...
func DBModelFilterExtendMap(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {

	if len(option.ExtendMap) > 0 {
		for key, value := range option.ExtendMap {
			dbModel = dbModel.Where(fmt.Sprintf("%v = %v", key, value))
		}
	}

	return dbModel, nil
}

//FieldMap ...
func FieldMap(exported string, extendMap map[string]string) string {
	if extendMap != nil {
		result, ok := extendMap[exported]
		if ok {
			return result
		}
	}
	result, ok := fieldMap[exported]
	if ok {
		return result
	}
	return ""
}

var (
	// fieldMap 查询字段映射
	fieldMap = map[string]string{
		"id":              "id",
		"name":            "name",
		"createTimestamp": "create_timestamp",
		"title":           "title",
		"creator":         "creator",
		"modifyTimestamp": "modify_time_stamp",
		"modifier":        "modifier",
		"deleteTimestamp": "delete_timestamp",
	}
)
