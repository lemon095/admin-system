// 自动生成模板Bullet
package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 子弹配置 结构体  Bullet
type Bullet struct {
	global.GVA_MODEL
	Name  *string        `json:"name" form:"name" gorm:"column:name;size:8;"`                  //name
	Value datatypes.JSON `json:"value" form:"value" gorm:"column:value;" swaggertype:"object"` //value
}

// TableName 子弹配置 Bullet自定义表名 configure_bullet
func (Bullet) TableName() string {
	return "configure_bullet"
}
