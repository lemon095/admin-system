package models

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetOperator(string)
	Generate() ActiveRecord
	GetId() interface{}
}
