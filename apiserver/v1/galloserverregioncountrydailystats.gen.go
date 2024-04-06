// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloServerRegionCountryDailyStat = "galloServerRegionCountryDailyStats"

// GalloServerRegionCountryDailyStat 服务用户区域分布统计（按天）
type GalloServerRegionCountryDailyStat struct {
	ID                  int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`       // ID
	ServerID            int32  `gorm:"column:serverId;comment:服务ID" json:"serverId"`                       // 服务ID
	CountryID           int32  `gorm:"column:countryId;comment:国家/区域ID" json:"countryId"`                  // 国家/区域ID
	Day                 string `gorm:"column:day;comment:日期YYYYMMDD" json:"day"`                           // 日期YYYYMMDD
	CountRequests       int64  `gorm:"column:countRequests;comment:请求数量" json:"countRequests"`             // 请求数量
	CountAttackRequests int64  `gorm:"column:countAttackRequests;comment:攻击数量" json:"countAttackRequests"` // 攻击数量
	AttackBytes         int64  `gorm:"column:attackBytes;comment:攻击流量" json:"attackBytes"`                 // 攻击流量
	Bytes               int64  `gorm:"column:bytes;comment:总流量" json:"bytes"`                              // 总流量
}

// TableName GalloServerRegionCountryDailyStat's table name
func (*GalloServerRegionCountryDailyStat) TableName() string {
	return TableNameGalloServerRegionCountryDailyStat
}
