// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloNodeAction = "galloNodeActions"

// GalloNodeAction 节点智能调度设置
type GalloNodeAction struct {
	ID       int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	NodeID   int64  `gorm:"column:nodeId;comment:节点ID" json:"nodeId"`                     // 节点ID
	Role     string `gorm:"column:role;comment:角色" json:"role"`                           // 角色
	IsOn     bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	Conds    string `gorm:"column:conds;comment:条件" json:"conds"`                         // 条件
	Action   string `gorm:"column:action;comment:动作" json:"action"`                       // 动作
	Duration string `gorm:"column:duration;comment:持续时间" json:"duration"`                 // 持续时间
	Order    int32  `gorm:"column:order;comment:排序" json:"order"`                         // 排序
	State    bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloNodeAction's table name
func (*GalloNodeAction) TableName() string {
	return TableNameGalloNodeAction
}
