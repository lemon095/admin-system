// 自动生成模板Weapon
package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 武器配置 结构体  Weapon
type Weapon struct {
	global.GVA_MODEL
	Name  *string        `json:"name" form:"name" gorm:"column:name;size:8;"`                  //name
	Value datatypes.JSON `json:"value" form:"value" gorm:"column:value;" swaggertype:"object"` //value
}

// TableName 武器配置 Weapon自定义表名 configure_weapon
func (Weapon) TableName() string {
	return "configure_weapon"
}
