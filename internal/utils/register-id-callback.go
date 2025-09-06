package utils

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterIDCallback[T any](db *gorm.DB, prefix string) {
	db.Callback().Create().Before("gorm:create").Register(fmt.Sprintf("set_id_%T", *new(T)), func(tx *gorm.DB) {
		stmt := tx.Statement
		if stmt.Schema != nil {
			field := stmt.Schema.LookUpField("ID")
			if field != nil {
				// fieldValue obtiene el valor del campo ID
				fieldValue, _ := field.ValueOf(tx.Statement.Context, stmt.ReflectValue)
				if fieldValue == "" {
					newID := fmt.Sprintf("%s-%s", prefix, uuid.New().String())
					_ = field.Set(tx.Statement.Context, stmt.ReflectValue, newID)
				}
			}
		}
	})
}
