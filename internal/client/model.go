package client

import "time"

type Client struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ClientID     string    `gorm:"uniqueIndex;not null" json:"client_id"`
	Secret       string    `gorm:"not null" json:"-"`
	Name         string    `json:"name"`
	Description  string    `gorm:"type:text" json:"description,omitempty"`
	RedirectURIs []string  `gorm:"-" json:"redirect_uris,omitempty"` // si usas otro modelo puedes serializar en jsonb
	RealmID      *string   `gorm:"type:uuid" json:"realm_id,omitempty"`
	Active       bool      `gorm:"default:true" json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
