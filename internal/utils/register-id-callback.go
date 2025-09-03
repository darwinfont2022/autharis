package utils

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterIDCallback(db *gorm.DB, model any, prefix string) {
	db.Callback().Create().Before("gorm:create").Register(fmt.Sprintf("set_id_%T", model), func(tx *gorm.DB) {
		if tx.Statement.Schema != nil && tx.Statement.Schema.ModelType == tx.Statement.ReflectValue.Type() {
			field := tx.Statement.Schema.LookUpField("ID")
			if field != nil {
				v, isZero := field.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue)
				if isZero {
					newID := fmt.Sprintf("%s-%s", prefix, uuid.New().String())
					_ = field.Set(tx.Statement.Context, tx.Statement.ReflectValue, newID)
				}
			}
		}
	})
}
