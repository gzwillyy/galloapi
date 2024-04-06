// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloNSPlan = "galloNSPlans"

// GalloNSPlan NS套餐
type GalloNSPlan struct {
	ID           int32   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"`     // ID
	Name         string  `gorm:"column:name;comment:套餐名称" json:"name"`                             // 套餐名称
	IsOn         bool    `gorm:"column:isOn;comment:是否启用" json:"isOn"`                             // 是否启用
	MonthlyPrice float64 `gorm:"column:monthlyPrice;default:0.00;comment:月价格" json:"monthlyPrice"` // 月价格
	YearlyPrice  float64 `gorm:"column:yearlyPrice;default:0.00;comment:年价格" json:"yearlyPrice"`   // 年价格
	Config       string  `gorm:"column:config;comment:配置" json:"config"`                           // 配置
	Order        int32   `gorm:"column:order;comment:排序" json:"order"`                             // 排序
	State        bool    `gorm:"column:state;default:1;comment:状态" json:"state"`                   // 状态
}

// TableName GalloNSPlan's table name
func (*GalloNSPlan) TableName() string {
	return TableNameGalloNSPlan
}