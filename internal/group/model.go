package group

import (
	"time"

	"autharis/internal/role"
)

type Group struct {
	ID      string  `gorm:"type:string;primaryKey" json:"id"`
	RealmID *string `gorm:"type:string" json:"realm_id,omitempty"`

	Name  string       `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Roles []*role.Role `gorm:"many2many:group_roles;constraint:OnDelete:CASCADE" json:"roles,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Pivot tables (GORM las maneja solo con many2many, pero las dejamos explícitas por control extra)
type UserGroup struct {
	UserID  string `gorm:"type:string;primaryKey"`
	GroupID string `gorm:"type:string;primaryKey"`
}

type GroupRole struct {
	UserID  string `gorm:"type:string;primaryKey"`
	GroupID string `gorm:"type:string;primaryKey"`
}
