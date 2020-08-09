package option

//DeleteOption 通用 Delete option
type DeleteOption struct {
	// ExtendMap 拓展字段
	ExtendMap map[string]string `json:"extendMap,omitempty"`
}

//AddExtendMapKV ...
func (o *DeleteOption) setExtendMapKV(k, v string) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]string)
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
func SetExtendMapDeleteOption(key, value string) DeleteOpt {
	return func(o *DeleteOption) *DeleteOption {
		o.setExtendMapKV(key, value)
		return o
	}
}
