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
	ExtendMap map[string]string `json:"extendMap,omitempty"`
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

func (o *ListOption) addExtendMapKV(k, v string) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]string)
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
func SetExtendMapListOption(key, value string) ListOpt {
	return func(o *ListOption) *ListOption {
		o.addExtendMapKV(key, value)
		return o
	}
}
