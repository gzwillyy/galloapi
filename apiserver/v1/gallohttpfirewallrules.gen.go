// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloHTTPFirewallRule = "galloHTTPFirewallRules"

// GalloHTTPFirewallRule 防火墙规则
type GalloHTTPFirewallRule struct {
	ID                int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	IsOn              bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                               // 是否启用
	Description       string `gorm:"column:description;comment:说明" json:"description"`                             // 说明
	Param             string `gorm:"column:param;comment:参数" json:"param"`                                         // 参数
	ParamFilters      string `gorm:"column:paramFilters;comment:处理器" json:"paramFilters"`                          // 处理器
	Operator          string `gorm:"column:operator;comment:操作符" json:"operator"`                                  // 操作符
	Value             string `gorm:"column:value;comment:对比值" json:"value"`                                        // 对比值
	IsCaseInsensitive bool   `gorm:"column:isCaseInsensitive;default:1;comment:是否大小写不敏感" json:"isCaseInsensitive"` // 是否大小写不敏感
	CheckpointOptions string `gorm:"column:checkpointOptions;comment:检查点参数" json:"checkpointOptions"`              // 检查点参数
	State             bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                               // 状态
	CreatedAt         int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                               // 创建时间
	AdminID           int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                                  // 管理员ID
	UserID            int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                                     // 用户ID
}

// TableName GalloHTTPFirewallRule's table name
func (*GalloHTTPFirewallRule) TableName() string {
	return TableNameGalloHTTPFirewallRule
}