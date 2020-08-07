package option

//ListOption list操作通用配置项
type ListOption struct {

	// Page 分页相关
	Page      Page              `json:"page,omitempty"`

	// Sort 排序 eg: []string{"id desc,create_time asc", "update_time"}
	Sort      string            `json:"sort,omitempty"`

	// Filter 只返回指定的字段，filter 为空则全部返回
	Filter    []string          `json:"filter,omitempty"`

	// ExtendMap 拓展字段
	ExtendMap map[string]string `json:"extendMap,omitempty"`
}

//Page ...
type Page struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
