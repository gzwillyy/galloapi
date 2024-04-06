// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloADPackageInstance = "galloADPackageInstances"

// GalloADPackageInstance 高防实例
type GalloADPackageInstance struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	IsOn           bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`               // 是否启用
	PackageID      int32  `gorm:"column:packageId;comment:规格ID" json:"packageId"`               // 规格ID
	ClusterID      int32  `gorm:"column:clusterId;comment:集群ID" json:"clusterId"`               // 集群ID
	NodeIds        string `gorm:"column:nodeIds;comment:节点ID" json:"nodeIds"`                   // 节点ID
	IPAddresses    string `gorm:"column:ipAddresses;comment:IP地址" json:"ipAddresses"`           // IP地址
	UserID         int64  `gorm:"column:userId;comment:用户ID" json:"userId"`                     // 用户ID
	UserDayTo      string `gorm:"column:userDayTo;comment:用户有效期YYYYMMDD" json:"userDayTo"`      // 用户有效期YYYYMMDD
	UserInstanceID int64  `gorm:"column:userInstanceId;comment:用户实例ID" json:"userInstanceId"`   // 用户实例ID
	State          bool   `gorm:"column:state;default:1;comment:状态" json:"state"`               // 状态
	ObjectCodes    string `gorm:"column:objectCodes;comment:防护对象" json:"objectCodes"`           // 防护对象
}

// TableName GalloADPackageInstance's table name
func (*GalloADPackageInstance) TableName() string {
	return TableNameGalloADPackageInstance
}
