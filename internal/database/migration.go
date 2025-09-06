package database

import (
	"autharis/internal/group"
	"autharis/internal/permission"
	"autharis/internal/realm"
	"autharis/internal/role"
	"autharis/internal/user"
	"autharis/internal/utils"
	"log"
)

func InitMigration() {
	// Registrar callbacks de IDs personalizados
	utils.RegisterIDCallback(DB, &realm.Realm{}, "RL")
	utils.RegisterIDCallback(DB, &role.Role{}, "RO")
	utils.RegisterIDCallback(DB, &permission.Permission{}, "PM")
	utils.RegisterIDCallback(DB, &user.User{}, "US")
	utils.RegisterIDCallback(DB, &group.Group{}, "GR")

	// Migraciones automáticas
	if err := DB.AutoMigrate(
		&realm.Realm{},
		&role.Role{},
		&permission.Permission{},
		&user.User{},
		&group.Group{},
		&group.UserGroup{},
		&group.GroupRole{},
	); err != nil {
		log.Fatalf("❌ Error en migraciones: %v", err)
	}

	log.Println("✅ Migraciones aplicadas correctamente")
}
