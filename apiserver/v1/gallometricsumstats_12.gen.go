// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloMetricSumStats12 = "galloMetricSumStats_12"

// GalloMetricSumStats12 指标统计总和数据
type GalloMetricSumStats12 struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	ClusterID  int32   `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`               // 集群ID
	NodeID     int32   `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                     // 节点ID
	ServerID   int32   `gorm:"column:serverId;comment:服务ID" json:"serverId"`                 // 服务ID
	ItemID     int64   `gorm:"column:itemId;comment:指标" json:"itemId"`                       // 指标
	Count      int64   `gorm:"column:count;comment:数量" json:"count"`                         // 数量
	Total      float64 `gorm:"column:total;default:0.00;comment:总和" json:"total"`            // 总和
	Time       string  `gorm:"column:time;comment:分钟值YYYYMMDDHHII" json:"time"`              // 分钟值YYYYMMDDHHII
	Version    int32   `gorm:"column:version;comment:版本号" json:"version"`                    // 版本号
	CreatedDay string  `gorm:"column:createdDay;comment:创建日期YYYYMMDD" json:"createdDay"`     // 创建日期YYYYMMDD
}

// TableName GalloMetricSumStats12's table name
func (*GalloMetricSumStats12) TableName() string {
	return TableNameGalloMetricSumStats12
}
