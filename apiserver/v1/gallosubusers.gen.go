// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloSubUser = "galloSubUsers"

// GalloSubUser 子用户
type GalloSubUser struct {
	ID       int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	UserID   int32  `gorm:"column:userId;comment:所属主用户ID" json:"userId"`                  // 所属主用户ID
	IsOn     bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	Name     string `gorm:"column:name;comment:名称" json:"name"`                           // 名称
	Username string `gorm:"column:username;comment:用户名" json:"username"`                  // 用户名
	Password string `gorm:"column:password;comment:密码" json:"password"`                   // 密码
	State    bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
}

// TableName GalloSubUser's table name
func (*GalloSubUser) TableName() string {
	return TableNameGalloSubUser
}