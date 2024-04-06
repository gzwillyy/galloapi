// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloReportResult = "galloReportResults"

// GalloReportResult 连通性监控结果
type GalloReportResult struct {
	ID           int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Type         string  `gorm:"column:type;comment:对象类型" json:"type"`                         // 对象类型
	TargetID     int64   `gorm:"column:targetId;comment:对象ID" json:"targetId"`                 // 对象ID
	TargetDesc   string  `gorm:"column:targetDesc;comment:对象描述" json:"targetDesc"`             // 对象描述
	UpdatedAt    int64   `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt"`               // 更新时间
	ReportNodeID int32   `gorm:"column:reportNodeId;comment:监控节点ID" json:"reportNodeId"`       // 监控节点ID
	IsOk         bool    `gorm:"column:isOk;default:1;comment:是否可连接" json:"isOk"`              // 是否可连接
	Level        string  `gorm:"column:level;comment:级别" json:"level"`                         // 级别
	CostMs       float64 `gorm:"column:costMs;default:0.00;comment:单次连接花费的时间" json:"costMs"`   // 单次连接花费的时间
	Error        string  `gorm:"column:error;comment:产生的错误信息" json:"error"`                    // 产生的错误信息
	CountUp      int32   `gorm:"column:countUp;comment:连续上线次数" json:"countUp"`                 // 连续上线次数
	CountDown    int32   `gorm:"column:countDown;comment:连续下线次数" json:"countDown"`             // 连续下线次数
}

// TableName GalloReportResult's table name
func (*GalloReportResult) TableName() string {
	return TableNameGalloReportResult
}
