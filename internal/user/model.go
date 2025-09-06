package user

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID           string `gorm:"type:string;primaryKey" json:"id"`
	Username     string `gorm:"uniqueIndex;not null" json:"username"`
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Active       bool   `gorm:"default:true" json:"active"`
	System       bool   `gorm:"default:false" json:"system"`

	Attributes datatypes.JSONMap `gorm:"type:jsonb" json:"attributes,omitempty"`

	RealmID *string `gorm:"type:string" json:"realm_id,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
