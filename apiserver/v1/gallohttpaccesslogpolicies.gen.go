// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloHTTPAccessLogPolicy = "galloHTTPAccessLogPolicies"

// GalloHTTPAccessLogPolicy 访问日志策略
type GalloHTTPAccessLogPolicy struct {
	ID               int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`        // ID
	TemplateID       int32  `gorm:"column:templateId;comment:模版ID" json:"templateId"`                    // 模版ID
	AdminID          int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                         // 管理员ID
	UserID           int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                            // 用户ID
	State            bool   `gorm:"column:state;default:1;comment:状态" json:"state"`                      // 状态
	CreatedAt        int64  `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                      // 创建时间
	Name             string `gorm:"column:name;comment:名称" json:"name"`                                  // 名称
	IsOn             bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                      // 是否启用
	Type             string `gorm:"column:type;comment:存储类型" json:"type"`                                // 存储类型
	Options          string `gorm:"column:options;comment:存储选项" json:"options"`                          // 存储选项
	Conds            string `gorm:"column:conds;comment:请求条件" json:"conds"`                              // 请求条件
	IsPublic         bool   `gorm:"column:isPublic;comment:是否为公用" json:"isPublic"`                       // 是否为公用
	FirewallOnly     bool   `gorm:"column:firewallOnly;comment:是否只记录防火墙相关" json:"firewallOnly"`          // 是否只记录防火墙相关
	Version          int32  `gorm:"column:version;comment:版本号" json:"version"`                           // 版本号
	DisableDefaultDB bool   `gorm:"column:disableDefaultDB;comment:是否停止默认数据库存储" json:"disableDefaultDB"` // 是否停止默认数据库存储
}

// TableName GalloHTTPAccessLogPolicy's table name
func (*GalloHTTPAccessLogPolicy) TableName() string {
	return TableNameGalloHTTPAccessLogPolicy
}