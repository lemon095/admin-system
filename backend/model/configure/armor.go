// 自动生成模板Armor
package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 防具配置 结构体  Armor
type Armor struct {
	global.GVA_MODEL
	Name  *string        `json:"name" form:"name" gorm:"column:name;size:8;"`                  //name
	Value datatypes.JSON `json:"value" form:"value" gorm:"column:value;" swaggertype:"object"` //value
}

// TableName 防具配置 Armor自定义表名 configure_armor
func (Armor) TableName() string {
	return "configure_armor"
}
