// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloLogin = "galloLogins"

// GalloLogin 第三方登录认证
type GalloLogin struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	AdminID int32  `gorm:"column:adminId;comment:管理员ID" json:"adminId"`                  // 管理员ID
	UserID  int32  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	IsOn    bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	Type    string `gorm:"column:type;comment:认证方式" json:"type"`                         // 认证方式
	Params  string `gorm:"column:params;comment:参数" json:"params"`                       // 参数
	State   bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloLogin's table name
func (*GalloLogin) TableName() string {
	return TableNameGalloLogin
}
