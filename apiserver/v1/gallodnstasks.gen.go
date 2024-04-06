// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloDNSTask = "galloDNSTasks"

// GalloDNSTask DNS更新任务
type GalloDNSTask struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	ClusterID  int32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`               // 集群ID
	ServerID   int32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`                 // 服务ID
	NodeID     int32  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                     // 节点ID
	DomainID   int32  `gorm:"column:domainId;comment:域名ID" json:"domainId"`                 // 域名ID
	RecordName string `gorm:"column:recordName;comment:记录名" json:"recordName"`              // 记录名
	Type       string `gorm:"column:type;comment:任务类型" json:"type"`                         // 任务类型
	UpdatedAt  int64  `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt"`               // 更新时间
	IsDone     bool   `gorm:"column:isDone;comment:是否已完成" json:"isDone"`                    // 是否已完成
	IsOk       bool   `gorm:"column:isOk;comment:是否成功" json:"isOk"`                         // 是否成功
	Error      string `gorm:"column:error;comment:错误信息" json:"error"`                       // 错误信息
	Version    int64  `gorm:"column:version;comment:版本" json:"version"`                     // 版本
	CountFails int32  `gorm:"column:countFails;comment:尝试失败次数" json:"countFails"`           // 尝试失败次数
}

// TableName GalloDNSTask's table name
func (*GalloDNSTask) TableName() string {
	return TableNameGalloDNSTask
}
