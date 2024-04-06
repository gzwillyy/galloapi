// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloMetricChart = "galloMetricCharts"

// GalloMetricChart 指标图表
type GalloMetricChart struct {
	ID              int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`           // ID
	ItemID          int32  `gorm:"column:itemId;comment:指标ID" json:"itemId"`                               // 指标ID
	Name            string `gorm:"column:name;comment:名称" json:"name"`                                     // 名称
	Code            string `gorm:"column:code;comment:代号" json:"code"`                                     // 代号
	Type            string `gorm:"column:type;comment:图形类型" json:"type"`                                   // 图形类型
	WidthDiv        int32  `gorm:"column:widthDiv;comment:宽度划分" json:"widthDiv"`                           // 宽度划分
	Params          string `gorm:"column:params;comment:图形参数" json:"params"`                               // 图形参数
	Order           int32  `gorm:"column:order;comment:排序" json:"order"`                                   // 排序
	IsOn            bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                         // 是否启用
	State           bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                         // 状态
	MaxItems        int32  `gorm:"column:maxItems;comment:最多条目" json:"maxItems"`                           // 最多条目
	IgnoreEmptyKeys bool   `gorm:"column:ignoreEmptyKeys;default:1;comment:忽略空的键值" json:"ignoreEmptyKeys"` // 忽略空的键值
	IgnoredKeys     string `gorm:"column:ignoredKeys;comment:忽略键值" json:"ignoredKeys"`                     // 忽略键值
}

// TableName GalloMetricChart's table name
func (*GalloMetricChart) TableName() string {
	return TableNameGalloMetricChart
}
