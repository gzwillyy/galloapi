// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloUserOrder = "galloUserOrders"

// GalloUserOrder 用户订单
type GalloUserOrder struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户订单" json:"id"` // 用户订单
	UserID      int64   `gorm:"column:userId;comment:用户ID" json:"userId"`                       // 用户ID
	Code        string  `gorm:"column:code;comment:订单号" json:"code"`                            // 订单号
	Type        string  `gorm:"column:type;comment:订单类型" json:"type"`                           // 订单类型
	MethodID    int32   `gorm:"column:methodId;comment:支付方式" json:"methodId"`                   // 支付方式
	Status      string  `gorm:"column:status;comment:订单状态" json:"status"`                       // 订单状态
	Amount      float64 `gorm:"column:amount;default:0.00;comment:金额" json:"amount"`            // 金额
	Params      string  `gorm:"column:params;comment:附加参数" json:"params"`                       // 附加参数
	ExpiredAt   int64   `gorm:"column:expiredAt;comment:过期时间" json:"expiredAt"`                 // 过期时间
	CreatedAt   int64   `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                 // 创建时间
	CancelledAt int64   `gorm:"column:cancelledAt;comment:取消时间" json:"cancelledAt"`             // 取消时间
	FinishedAt  int64   `gorm:"column:finishedAt;comment:结束时间" json:"finishedAt"`               // 结束时间
	State       bool    `gorm:"column:state;default:1;comment:状态" json:"state"`                 // 状态
}

// TableName GalloUserOrder's table name
func (*GalloUserOrder) TableName() string {
	return TableNameGalloUserOrder
}
