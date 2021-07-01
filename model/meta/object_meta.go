package meta

//ObjectMeta 对象 metadata
type ObjectMeta struct {

	// ID 唯一标识
	ID ID `json:"id" db:"id" bson:"id" gorm:"column:id"`

	// Name 英文名 如果对象是User name 为用户的唯一字符串标识 如 leoshli
	Name Name `json:"name" db:"name" bson:"name" gorm:"column:name"`

	// Title 资源中文名 如果对象是User 则Title 例： 李少辉
	Title string `json:"title" db:"title" bson:"title" gorm:"column:title"`

	// CreateTimestamp 资源创建时间戳
	CreateTimestamp int64 `json:"createTimestamp" db:"create_timestamp" bson:"create_timestamp" gorm:"column:create_timestamp"`

	// Creator 创建人Name
	Creator string `json:"creator" db:"creator" bson:"creator" gorm:"column:creator"`

	// ModifyTimeStamp 资源更新时间戳
	ModifyTimestamp int64 `json:"modifyTimestamp" db:"modify_timestamp" bson:"modify_timestamp"gorm:"column:modify_timestamp"`

	// Modifier 更新人Name
	Modifier string `json:"modifier" db:"modifier" bson:"modifier" gorm:"column:modifier"`

	// DeleteTimestamp 删除时间，零值标识未删除
	DeleteTimestamp int64 `json:"deleteTimestamp" db:"delete_timestamp" bson:"delete_timestamp" gorm:"column:delete_timestamp"`

	// Isolation 资源隔离标识 默认隔离业务ID appID 例如 {"5608":""}
	Isolation map[string]string `json:"isolation" db:"isolation" bson:"isolation" gorm:"column:isolation"`

	// Labels 资源标签标识
	Labels map[string]string `json:"labels" db:"labels" bson:"labels" gorm:"column:labels"`

	// Annotations 附加信息
	Annotations map[string]string `json:"annotations" db:"annotations" bson:"annotations" gorm:"column:annotations"`
}

//ListMeta ...
type ListMeta struct {
	ObjectMeta `json:",inline" bson:",inline"`
	Total      int64 `json:"total"`
}
