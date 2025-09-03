package user

import (
	"time"

	"github.com/darwinfont2022/autharis/internal/role"
)

type User struct {
	ID           string            `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username     string            `gorm:"uniqueIndex;not null" json:"username"`
	Email        string            `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string            `gorm:"not null" json:"-"`
	FirstName    string            `json:"first_name,omitempty"`
	LastName     string            `json:"last_name,omitempty"`
	Active       bool              `gorm:"default:true" json:"active"`
	System       bool              `gorm:"default:false" json:"system"`
	Attributes   map[string]string `gorm:"-" json:"attributes,omitempty"` // atributos customizados
	Roles        []*role.Role      `gorm:"many2many:user_roles;constraint:OnDelete:CASCADE" json:"roles,omitempty"`
	RealmID      *string           `gorm:"type:uuid" json:"realm_id,omitempty"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
