package option

import (
	"fmt"
	"strings"
	"unicode"

	"git.code.oa.com/trpcprotocol/tab/common"
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
	Filter map[string]*common.FilterItem `json:"filter,omitempty"`

	// ExtendMap 拓展字段 保留
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
func WithFilterListOption(filter map[string]*common.FilterItem) ListOpt {
	return func(o *ListOption) *ListOption {
		o.Filter = filter
		return o
	}
}

//WithFilterEntryListOption 设置单个查询条件
func WithFilterEntryListOption(key string, value *common.FilterItem) ListOpt {
	return func(o *ListOption) *ListOption {
		if o.Filter == nil {
			o.Filter = make(map[string]*common.FilterItem)
		}
		o.Filter[key] = value
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
		DBModelFilter,
		DBModelFilterOrderBy,
		DBModelFilterPage,
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

	if len(option.Filter) > 0 {
		for key, value := range option.Filter {
			if len(key) == 0 {
				continue
			}
			key = canonicalKey(key)
			if len(value.Value) == 0 {
				dbModel = dbModel.Where(fmt.Sprintf(`%v %v ""`, key, operator(value.Operator)))
				continue
			}
			switch value.Operator {
			case common.Operator_Between:
				valueList := strings.Split(value.Value, ";")
				if len(valueList) != 2 { // 错误的数据
					return dbModel, errors.Errorf("invalid between-value:%v", value.Value)
				}
				dbModel = dbModel.Where(fmt.Sprintf(`%v <= %v and %v <= %v`, valueList[0], key, key, valueList[1]))
			default:
				dbModel = dbModel.Where(fmt.Sprintf(`%v %v %v`, key, operator(value.Operator), value.Value))
			}
		}
	}
	return dbModel, nil
}

func operator(o common.Operator) string {
	switch o {
	case common.Operator_EQUAL:
		return "="
	case common.Operator_LT:
		return "<"
	case common.Operator_LTE:
		return "<="
	case common.Operator_GT:
		return ">"
	case common.Operator_GTE:
		return ">="
	case common.Operator_IN:
		return "in"
	case common.Operator_NOT_IN:
		return "not in"
	case common.Operator_LIKE:
		return "like"
	default:
		return "="
	}
}

func canonicalKey(key string) string {
	return underscoreName(key)
}

func underscoreName(name string) string {
	buffer := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteString(strings.ToLower(string(r)))
		} else {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
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
