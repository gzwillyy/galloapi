// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserTrafficBill = "galloUserTrafficBills"

// GalloUserTrafficBill 用户流量/带宽账单
type GalloUserTrafficBill struct {
	ID                    int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`                        // ID
	BillID                int64   `gorm:"column:billId;not null;comment:主账单ID" json:"billId"`                                  // 主账单ID
	RegionID              int32   `gorm:"column:regionId;comment:区域ID" json:"regionId"`                                        // 区域ID
	Amount                float64 `gorm:"column:amount;default:0.00;comment:金额" json:"amount"`                                 // 金额
	BandwidthMB           float64 `gorm:"column:bandwidthMB;default:0.0000;comment:带宽MB" json:"bandwidthMB"`                   // 带宽MB
	BandwidthPercentile   int32   `gorm:"column:bandwidthPercentile;comment:带宽百分位" json:"bandwidthPercentile"`                 // 带宽百分位
	TrafficGB             float64 `gorm:"column:trafficGB;default:0.00000000;comment:流量GB" json:"trafficGB"`                   // 流量GB
	TrafficPackageGB      float64 `gorm:"column:trafficPackageGB;default:0.00000000;comment:使用的流量包GB" json:"trafficPackageGB"` // 使用的流量包GB
	UserTrafficPackageIds string  `gorm:"column:userTrafficPackageIds;comment:使用的流量包ID" json:"userTrafficPackageIds"`          // 使用的流量包ID
	PricePerUnit          float64 `gorm:"column:pricePerUnit;default:0.0000;comment:单位价格" json:"pricePerUnit"`                 // 单位价格
	PriceType             string  `gorm:"column:priceType;comment:计费方式：traffic|bandwidth" json:"priceType"`                    // 计费方式：traffic|bandwidth
	State                 bool    `gorm:"column:state;default:1;comment:状态" json:"state"`                                      // 状态
}

// TableName GalloUserTrafficBill's table name
func (*GalloUserTrafficBill) TableName() string {
	return TableNameGalloUserTrafficBill
}
