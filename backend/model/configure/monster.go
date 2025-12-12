// 自动生成模板Monster
package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 怪物配置 结构体  Monster
type Monster struct {
	global.GVA_MODEL
	Name  *string        `json:"name" form:"name" gorm:"column:name;size:8;"`                  //name
	Value datatypes.JSON `json:"value" form:"value" gorm:"column:value;" swaggertype:"object"` //value
}

// TableName 怪物配置 Monster自定义表名 configure_monster
func (Monster) TableName() string {
	return "configure_monster"
}
