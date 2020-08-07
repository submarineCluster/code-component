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

//SetLimit ...
func (o *ListOption) SetLimit(limit int64) {
	o.Page.Limit = limit
}

//SetOffset ...
func (o *ListOption) SetOffset(offset int64) {
	o.Page.Offset = offset
}

//AddExtendMapKV ...
func (o *ListOption) AddExtendMapKV(k, v string) {
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
