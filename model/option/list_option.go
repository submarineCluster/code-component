package option

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
