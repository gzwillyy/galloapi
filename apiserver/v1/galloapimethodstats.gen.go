// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloAPIMethodStat = "galloAPIMethodStats"

// GalloAPIMethodStat API方法统计
type GalloAPIMethodStat struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	APINodeID  int32   `gorm:"column:apiNodeId;comment:API节点ID" json:"apiNodeId"`            // API节点ID
	Method     string  `gorm:"column:method;comment:方法" json:"method"`                       // 方法
	Tag        string  `gorm:"column:tag;comment:标签方法" json:"tag"`                           // 标签方法
	CostMs     float64 `gorm:"column:costMs;default:0.00;comment:耗时Ms" json:"costMs"`        // 耗时Ms
	PeekMs     float64 `gorm:"column:peekMs;default:0.00;comment:峰值耗时" json:"peekMs"`        // 峰值耗时
	CountCalls int64   `gorm:"column:countCalls;comment:调用次数" json:"countCalls"`             // 调用次数
	Day        string  `gorm:"column:day;comment:日期" json:"day"`                             // 日期
}

// TableName GalloAPIMethodStat's table name
func (*GalloAPIMethodStat) TableName() string {
	return TableNameGalloAPIMethodStat
}
