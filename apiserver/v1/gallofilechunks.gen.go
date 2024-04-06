// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package v1

const TableNameGalloFileChunk = "galloFileChunks"

// GalloFileChunk 文件片段
type GalloFileChunk struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	FileID int32  `gorm:"column:fileId;comment:文件ID" json:"fileId"`                     // 文件ID
	Data   []byte `gorm:"column:data;comment:分块内容" json:"data"`                         // 分块内容
}

// TableName GalloFileChunk's table name
func (*GalloFileChunk) TableName() string {
	return TableNameGalloFileChunk
}
