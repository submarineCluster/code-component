package option

//GetOption 通用 Get option
type GetOption struct {
	//FlagBit 标志位
	FlagBit FlagBit `json:"flagBit,omitempty"`
	// ExtendMap 拓展字段
	ExtendMap map[string]string `json:"extendMap,omitempty"`
}

//FlagBit 标志位
type FlagBit uint16

const (
	//NormalFlagBit FlagBit = 0
	DeleteFlagBit FlagBit = 1 << iota // 包含被删除数据
)

//AddExtendMapKV ...
func (o *GetOption) AddExtendMapKV(k, v string) {
	if o.ExtendMap == nil {
		o.ExtendMap = make(map[string]string)
	}
	o.ExtendMap[k] = v
}

//GetOpt ...
type GetOpt func(o *GetOption) *GetOption

//WithDeleteGetOption ...
func WithDeleteGetOption() GetOpt {
	return func(o *GetOption) *GetOption {
		o.FlagBit |= DeleteFlagBit
		return o
	}
}
