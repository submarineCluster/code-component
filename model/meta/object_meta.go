package meta

//ObjectMeta 对象 metadata
type ObjectMeta struct {

	// ID 唯一标识
	ID ID `json:"id" db:"id" bson:"id" gorm:"column:id"`

	// Name 资源标识 eg: user, role, company etc
	Name Name `json:"name" db:"name" bson:"name" gorm:"column:name"`

	// Title 资源中文名 eg: 用户\角色\公司 等
	Title string `json:"title" db:"title" bson:"title" gorm:"column:title"`

	// CreateTimestamp 资源创建时间戳
	CreateTimestamp int64 `json:"createTimestamp" db:"create_timestamp" bson:"create_timestamp" gorm:"column:create_timestamp"`

	// Creator 创建人ID
	Creator int64 `json:"creator" db:"creator" bson:"creator" gorm:"column:creator"`

	// ModifyTimeStamp 资源更新时间戳
	ModifyTimeStamp int64 `json:"modifyTimeStamp" db:"modify_time_stamp" bson:"modify_time_stamp"gorm:"column:modify_time_stamp"`

	// Modifier 更新人ID
	Modifier int64 `json:"modifier" db:"modifier" bson:"modifier" gorm:"column:modifier"`

	// DeleteTimestamp 删除时间，零值标识未删除
	DeleteTimestamp int64 `json:"deleteTimestamp" db:"delete_timestamp" bson:"delete_timestamp" gorm:"column:delete_timestamp"`
}

//ListMeta ...
type ListMeta struct {
	ObjectMeta `json:",inline" bson:",inline"`
	Total      int64 `json:"total"`
}
