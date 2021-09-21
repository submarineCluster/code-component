package option

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//ListOption list操作通用配置项
type ListOption struct {

	//FlagBit 标志位
	FlagBit FlagBit `json:"flagBit,omitempty"`
	// Page 分页相关
	Page Page `json:"page,omitempty"`

	// OrderByList 排序
	OrderByList []OrderBy `json:"sort,omitempty"`

	// Filter 只返回指定的字段，Selector 为空则全部返回 TODO 还没支持
	Selector []string `json:"selector,omitempty"`

	// Filter 过滤条件
	Filter map[string]FilterItem `json:"filter,omitempty"`

	// ExtendMap 拓展字段 保留
	ExtendMap map[string]interface{} `json:"extendMap,omitempty"`
}

//Operator 操作符
type Operator string

// const ...
const (
	LtOperator    Operator = "<"
	LteOperator   Operator = "<="
	EqualOperator Operator = "="
	GtOperator    Operator = ">"
	GteOperator   Operator = ">="
	InOperator    Operator = "in"
	NotInOperator Operator = "not in"
	LikeOperator  Operator = "like"
)

//FilterItem ...
type FilterItem struct {
	Operator Operator `json:"operator"`
	Value    string   `json:"value"`
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

//WithLimitListOption ...
func WithLimitListOption(limit int64) ListOpt {
	return func(o *ListOption) *ListOption {
		o.setLimit(limit)
		return o
	}
}

//WithOffsetListOption ...
func WithOffsetListOption(offset int64) ListOpt {
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

//SetSelectorListOption 筛选字段
func SetSelectorListOption(key ...string) ListOpt {
	return func(o *ListOption) *ListOption {
		o.Selector = append(o.Selector, key...)
		return o
	}
}

//WithFilterListOption 设置查询条件
func WithFilterListOption(filter map[string]FilterItem) ListOpt {
	return func(o *ListOption) *ListOption {
		o.Filter = filter
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
		DBModelFilter,
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
		dbModel = dbModel.Limit(int(option.Page.Limit))
	}

	if option.Page.Offset >= 0 {
		dbModel = dbModel.Offset(int(option.Page.Offset))
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

//DBModelFilter ...
func DBModelFilter(dbModel *gorm.DB, option *ListOption) (*gorm.DB, error) {

	if len(option.ExtendMap) > 0 {
		for key, value := range option.Filter {
			dbModel = dbModel.Where(fmt.Sprintf("%v %v %v", key, value.Operator, value.Value))
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

//GetFieldMap ...
func GetFieldMap(extendMap map[string]string) map[string]string {
	result := map[string]string{}
	for key, value := range fieldMap {
		result[key] = value
	}
	for key, value := range extendMap {
		result[key] = value
	}
	return result
}

var (
	// fieldMap 查询字段映射
	fieldMap = map[string]string{
		"id":              "id",
		"name":            "name",
		"title":           "title",
		"creator":         "creator",
		"modifier":        "modifier",
		"modifyTimestamp": "modify_timestamp",
		"createTimestamp": "create_timestamp",
		"deleteTimestamp": "delete_timestamp",
	}
)
