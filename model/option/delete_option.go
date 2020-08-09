package option

//DeleteOption 通用 Delete option
type DeleteOption struct {
	// ExtendMap 拓展字段
	ExtendMap map[string]interface{} `json:"extendMap,omitempty"`
}

//AddExtendMapKV ...
func (o *DeleteOption) setExtendMapKV(k string, v interface{}) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]interface{})
	}
	o.ExtendMap[k] = v
}

//NewDeleteOption ...
func NewDeleteOption() *DeleteOption {
	return &DeleteOption{}
}

//DeleteOpt ...
type DeleteOpt func(o *DeleteOption) *DeleteOption

//SetExtendMapDeleteOption ...
func SetExtendMapDeleteOption(key string, value interface{}) DeleteOpt {
	return func(o *DeleteOption) *DeleteOption {
		o.setExtendMapKV(key, value)
		return o
	}
}
