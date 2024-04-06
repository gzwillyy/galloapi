// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserAccountDailyStat = "galloUserAccountDailyStats"

// GalloUserAccountDailyStat 账户每日统计
type GalloUserAccountDailyStat struct {
	ID      int32   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Day     string  `gorm:"column:day;comment:YYYYMMDD" json:"day"`                       // YYYYMMDD
	Month   string  `gorm:"column:month;comment:YYYYMM" json:"month"`                     // YYYYMM
	Income  float64 `gorm:"column:income;default:0.00;comment:收入" json:"income"`          // 收入
	Expense float64 `gorm:"column:expense;default:0.00;comment:支出" json:"expense"`        // 支出
}

// TableName GalloUserAccountDailyStat's table name
func (*GalloUserAccountDailyStat) TableName() string {
	return TableNameGalloUserAccountDailyStat
}
