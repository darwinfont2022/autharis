package permission

import "time"

type Permission struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Resource    string    `gorm:"not null" json:"resource"` // recurso sobre el que aplica
	Action      string    `gorm:"not null" json:"action"`   // acción (read, write, delete, etc.)
	RoleID      *string   `gorm:"type:uuid" json:"role_id,omitempty"`
	Active      bool      `gorm:"default:true" json:"active"`
	System      bool      `gorm:"default:false" json:"system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
