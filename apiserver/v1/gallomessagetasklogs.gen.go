// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloMessageTaskLog = "galloMessageTaskLogs"

// GalloMessageTaskLog 消息发送日志
type GalloMessageTaskLog struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	TaskID    int64  `gorm:"column:taskId;comment:任务ID" json:"taskId"`                     // 任务ID
	CreatedAt int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	IsOk      bool   `gorm:"column:isOk;comment:是否成功" json:"isOk"`                         // 是否成功
	Error     string `gorm:"column:error;comment:错误信息" json:"error"`                       // 错误信息
	Response  string `gorm:"column:response;comment:响应信息" json:"response"`                 // 响应信息
	Day       string `gorm:"column:day;comment:YYYYMMDD" json:"day"`                       // YYYYMMDD
}

// TableName GalloMessageTaskLog's table name
func (*GalloMessageTaskLog) TableName() string {
	return TableNameGalloMessageTaskLog
}