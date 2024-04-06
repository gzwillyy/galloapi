// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloADPackage = "galloADPackages"

// GalloADPackage 高防产品规格
type GalloADPackage struct {
	ID                      int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                 // ID
	IsOn                    bool   `gorm:"column:isOn;default:1;comment:是否启用" json:"isOn"`                               // 是否启用
	NetworkID               int32  `gorm:"column:networkId;comment:线路ID" json:"networkId"`                               // 线路ID
	ProtectionBandwidthSize int32  `gorm:"column:protectionBandwidthSize;comment:防护带宽尺寸" json:"protectionBandwidthSize"` // 防护带宽尺寸
	ProtectionBandwidthUnit string `gorm:"column:protectionBandwidthUnit;comment:防护带宽单位" json:"protectionBandwidthUnit"` // 防护带宽单位
	ProtectionBandwidthBits int64  `gorm:"column:protectionBandwidthBits;comment:防护带宽比特" json:"protectionBandwidthBits"` // 防护带宽比特
	ServerBandwidthSize     int32  `gorm:"column:serverBandwidthSize;comment:业务带宽尺寸" json:"serverBandwidthSize"`         // 业务带宽尺寸
	ServerBandwidthUnit     string `gorm:"column:serverBandwidthUnit;comment:业务带宽单位" json:"serverBandwidthUnit"`         // 业务带宽单位
	ServerBandwidthBits     int64  `gorm:"column:serverBandwidthBits;comment:业务带宽比特" json:"serverBandwidthBits"`         // 业务带宽比特
	State                   bool   `gorm:"column:state;comment:状态" json:"state"`                                         // 状态
}

// TableName GalloADPackage's table name
func (*GalloADPackage) TableName() string {
	return TableNameGalloADPackage
}