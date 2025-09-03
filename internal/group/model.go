package group

import (
	"Autharis/internal/role"
	"time"
)

type Group struct {
	ID        string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Roles     []*role.Role `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE" json:"roles,omitempty"`
	RealmID   []*string    `gorm:"type:uuid" json:"realm_id,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
