package option

//DeleteOption 通用 Delete option
type DeleteOption struct {
	// ExtendMap 拓展字段
	ExtendMap map[string]string `json:"extendMap,omitempty"`
}

//AddExtendMapKV ...
func (o *DeleteOption) AddExtendMapKV(k, v string) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]string)
	}
	o.ExtendMap[k] = v
}
