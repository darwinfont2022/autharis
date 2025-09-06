package database

import (
	"autharis/internal/permission"
	"autharis/internal/realm"
	"autharis/internal/role"
	"autharis/internal/user"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// SeedData crea datos iniciales si no existen
func SeedData() {
	// Crear Realm por defecto
	defaultRealm := realm.Realm{}
	if err := DB.Where("name = ?", "Master").First(&defaultRealm).Error; err != nil {
		defaultRealm = realm.Realm{Name: "Master"}
		if err := DB.Create(&defaultRealm).Error; err != nil {
			log.Fatalf("❌ Error creando Realm por defecto: %v", err)
		}
		log.Println("✅ Realm por defecto creado")
	}

	// Crear Role superadmin
	superAdminRole := role.Role{}
	if err := DB.Where("name = ? AND realm_id = ?", "superadmin", defaultRealm.ID).First(&superAdminRole).Error; err != nil {
		superAdminRole = role.Role{Name: "superadmin", RealmID: &defaultRealm.ID}
		if err := DB.Create(&superAdminRole).Error; err != nil {
			log.Fatalf("❌ Error creando Role superadmin: %v", err)
		}
		log.Println("✅ Role superadmin creado")
	}

	// Crear permisos básicos
	permissions := []permission.Permission{
		{
			Name:        "manage_realms",
			Description: "Crear, modificar y eliminar realms",
			Resource:    "realm",
			Action:      "manage",
			RealmID:     &defaultRealm.ID,
			Active:      true,
			System:      true,
		},
		{
			Name:        "manage_users",
			Description: "Crear, modificar y eliminar usuarios",
			Resource:    "user",
			Action:      "manage",
			RealmID:     &defaultRealm.ID,
			Active:      true,
			System:      true,
		},
		{
			Name:        "manage_roles",
			Description: "Gestionar roles y asignaciones",
			Resource:    "role",
			Action:      "manage",
			RealmID:     &defaultRealm.ID,
			Active:      true,
			System:      true,
		},
		{
			Name:        "manage_permissions",
			Description: "Gestionar permisos del sistema",
			Resource:    "permission",
			Action:      "manage",
			RealmID:     &defaultRealm.ID,
			Active:      true,
			System:      true,
		},
		{
			Name:        "manage_groups",
			Description: "Crear y modificar grupos de usuarios",
			Resource:    "group",
			Action:      "manage",
			RealmID:     &defaultRealm.ID,
			Active:      true,
			System:      true,
		},
	}
	var createdPerms []permission.Permission
	for _, p := range permissions {
		perm := permission.Permission{}

		// Buscar por nombre + realm_id solo si RealmID no es nil
		query := DB.Where("name = ?", p.Name)
		if p.RealmID != nil {
			query = query.Where("realm_id = ?", *p.RealmID)
		}

		err := query.First(&perm).Error
		if err != nil {
			// No encontrado → crear
			if err := DB.Create(&p).Error; err != nil {
				log.Fatalf("❌ Error creando Permission %s: %v", p.Name, err)
			}
			log.Printf("✅ Permission %s creado", p.Name)
			perm = p // asignar para agregar a la lista de permisos creados
		}

		createdPerms = append(createdPerms, perm)
	}

	// Relacionar Role superadmin con permisos
	if err := DB.Model(&superAdminRole).Association("Permissions").Replace(createdPerms); err != nil {
		log.Fatalf("❌ Error asignando permisos a superadmin: %v", err)
	}

	log.Println("✅ Permisos asignados al role superadmin")

	// Crear User admin
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminEmail == "" || adminPassword == "" {
		log.Println("⚠️ Variables ADMIN_EMAIL y ADMIN_PASSWORD no definidas, no se creó el admin")
		return
	}

	adminUser := user.User{}
	if err := DB.Where("email = ? AND realm_id = ?", adminEmail, defaultRealm.ID).First(&adminUser).Error; err != nil {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		adminUser = user.User{
			Username:     "admin",
			Email:        adminEmail,
			PasswordHash: string(hashed),
			RealmID:      &defaultRealm.ID,
		}
		if err := DB.Create(&adminUser).Error; err != nil {
			log.Fatalf("❌ Error creando User admin: %v", err)
		}
		log.Println("✅ User admin creado")
	}

	// Relacionar admin ↔ superadmin
	if err := DB.Model(&adminUser).Association("Roles").Append(&superAdminRole); err != nil {
		log.Fatalf("❌ Error asignando role superadmin al admin: %v", err)
	}
	log.Println("✅ Admin vinculado con superadmin")
}
