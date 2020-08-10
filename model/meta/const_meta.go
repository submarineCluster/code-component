package meta

//ObjectKind 对象定义唯一标识，eg user/role
type ObjectKind string

//APIVersion 对象api版本，拓展信息，不兼容的api改动时可标识 eg api/v1  api/v2
type APIVersion string

// String ...
func (k ObjectKind) String() string {
	return string(k)
}

// String ...
func (v APIVersion) String() string {
	return string(v)
}
