package role

import (
	"autharis/internal/permission"
	"time"

	"gorm.io/datatypes"
)

type Role struct {
	ID          string `gorm:"type:string;primaryKey" json:"id"`
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:text" json:"description,omitempty"`

	RealmID  *string `gorm:"type:string" json:"realm_id,omitempty"`
	ClientID *string `gorm:"type:string" json:"client_id,omitempty"`

	Composite bool    `gorm:"default:false" json:"composite"`
	Children  []*Role `gorm:"many2many:role_composites;constraint:OnDelete:CASCADE" json:"children,omitempty"`

	Attributes  datatypes.JSONMap        `gorm:"type:jsonb" json:"attributes,omitempty"`
	Permissions []*permission.Permission `gorm:"many2many:role_permissions"`
	Active      bool                     `gorm:"default:true" json:"active"`
	System      bool                     `gorm:"default:false" json:"system"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}
