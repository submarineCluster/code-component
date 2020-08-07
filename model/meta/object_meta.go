package meta

//ObjectMeta 对象 metadata
type ObjectMeta struct {

	// ID 唯一标识
	ID ID `json:"id" db:"id" bson:"id"`

	// Name 资源标识 eg: user, role, company etc
	Name Name `json:"name" db:"name" bson:"name"`

	// Title 资源中文名 eg: 用户\角色\公司 等
	Title string `json:"title" db:"title" bson:"title"`

	// CreateTimeStamp 资源创建时间戳
	CreateTimeStamp int64 `json:"createTimeStamp"`

	// Creator 创建人ID
	Creator int64 `json:"creator"`

	// ModifyTimeStamp 资源更新时间戳
	ModifyTimeStamp int64 `json:"modifyTimeStamp" db:"modify_time_stamp" bson:"modify_time_stamp"`

	// Modifier 更新人ID
	Modifier int64 `json:"modifier" db:"modifier" bson:"modifier"`
}

//ListMeta ...
type ListMeta struct {
	ObjectMeta `json:",inline" bson:",inline"`
	Total      int64 `json:"total"`
}
