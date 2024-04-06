// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloHTTPCacheTaskKey = "galloHTTPCacheTaskKeys"

// GalloHTTPCacheTaskKey 缓存任务Key
type GalloHTTPCacheTaskKey struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	TaskID    int64  `gorm:"column:taskId;comment:任务ID" json:"taskId"`                     // 任务ID
	Key       string `gorm:"column:key;comment:Key" json:"key"`                            // Key
	KeyType   string `gorm:"column:keyType;comment:Key类型：key|prefix" json:"keyType"`       // Key类型：key|prefix
	Type      string `gorm:"column:type;comment:操作类型" json:"type"`                         // 操作类型
	ClusterID int32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`               // 集群ID
	Nodes     string `gorm:"column:nodes;comment:节点" json:"nodes"`                         // 节点
	Errors    string `gorm:"column:errors;comment:错误信息" json:"errors"`                     // 错误信息
	IsDone    bool   `gorm:"column:isDone;comment:是否已完成" json:"isDone"`                    // 是否已完成
}

// TableName GalloHTTPCacheTaskKey's table name
func (*GalloHTTPCacheTaskKey) TableName() string {
	return TableNameGalloHTTPCacheTaskKey
}
