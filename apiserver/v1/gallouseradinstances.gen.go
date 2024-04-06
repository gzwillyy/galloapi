// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserADInstance = "galloUserADInstances"

// GalloUserADInstance 高防实例
type GalloUserADInstance struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	AdminID     int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
	UserID      int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	InstanceID  int32  `gorm:"column:instanceId;comment:高防实例ID" json:"instanceId"`           // 高防实例ID
	PeriodID    int32  `gorm:"column:periodId;comment:有效期" json:"periodId"`                  // 有效期
	PeriodCount int32  `gorm:"column:periodCount;comment:有效期数量" json:"periodCount"`          // 有效期数量
	PeriodUnit  string `gorm:"column:periodUnit;comment:有效期单位" json:"periodUnit"`            // 有效期单位
	DayFrom     string `gorm:"column:dayFrom;comment:开始日期" json:"dayFrom"`                   // 开始日期
	DayTo       string `gorm:"column:dayTo;comment:结束日期" json:"dayTo"`                       // 结束日期
	MaxObjects  int32  `gorm:"column:maxObjects;comment:最多防护对象数" json:"maxObjects"`          // 最多防护对象数
	ObjectCodes string `gorm:"column:objectCodes;comment:防护对象" json:"objectCodes"`           // 防护对象
	CreatedAt   int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloUserADInstance's table name
func (*GalloUserADInstance) TableName() string {
	return TableNameGalloUserADInstance
}
