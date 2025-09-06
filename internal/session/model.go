package session

import "time"

type Session struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	ClientID  *string   `gorm:"type:uuid" json:"client_id,omitempty"`
	RealmID   *string   `gorm:"type:uuid" json:"realm_id,omitempty"`
	Token     string    `gorm:"type:text;not null" json:"token"` // JWT o refresh token
	Active    bool      `gorm:"default:true" json:"active"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
