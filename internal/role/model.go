package role

import "time"

type Role struct {
	ID          string            `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string            `gorm:"uniqueIndex;not null" json:"name"`
	Description string            `gorm:"type:text" json:"description,omitempty"`
	RealmID     *string           `gorm:"type:uuid" json:"realm_id,omitempty"`
	ClientID    *string           `gorm:"type:uuid" json:"client_id,omitempty"`
	Composite   bool              `gorm:"default:false" json:"composite"`
	Children    []*Role           `gorm:"many2many:role_composites;constraint:OnDelete:CASCADE" json:"children,omitempty"`
	Attributes  map[string]string `gorm:"-" json:"attributes,omitempty"`
	Active      bool              `gorm:"default:true" json:"active"`
	System      bool              `gorm:"default:false" json:"system"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
