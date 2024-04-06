// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserNode = "galloUserNodes"

// GalloUserNode 用户节点
type GalloUserNode struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	IsOn        bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	UniqueID    string `gorm:"column:uniqueId;comment:唯一ID" json:"uniqueId"`                 // 唯一ID
	Secret      string `gorm:"column:secret;comment:密钥" json:"secret"`                       // 密钥
	Name        string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Description string `gorm:"column:description;comment:描述" json:"description"`             // 描述
	HTTP        string `gorm:"column:http;comment:监听的HTTP配置" json:"http"`                    // 监听的HTTP配置
	HTTPS       string `gorm:"column:https;comment:监听的HTTPS配置" json:"https"`                 // 监听的HTTPS配置
	AccessAddrs string `gorm:"column:accessAddrs;comment:外部访问地址" json:"accessAddrs"`         // 外部访问地址
	Order       int32  `gorm:"column:order;comment:排序" json:"order"`                         // 排序
	State       bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	CreatedAt   int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`               // 创建时间
	AdminID     int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
	Weight      int32  `gorm:"column:weight;comment:权重" json:"weight"`                       // 权重
	Status      string `gorm:"column:status;comment:运行状态" json:"status"`                     // 运行状态
}

// TableName GalloUserNode's table name
func (*GalloUserNode) TableName() string {
	return TableNameGalloUserNode
}