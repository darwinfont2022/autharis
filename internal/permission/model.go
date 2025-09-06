package permission

import "time"

type Permission struct {
	ID          string    `gorm:"type:string;primaryKey" json:"id"`
	Name        string    `gorm:"size:150;not null;index:idx_perm_name_realm,unique" json:"name"`
	RealmID     *string   `gorm:"type:string;index:idx_perm_name_realm,unique" json:"realm_id,omitempty"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Resource    string    `gorm:"size:150;not null" json:"resource"`
	Action      string    `gorm:"size:50;not null" json:"action"`
	RoleID      *string   `gorm:"type:string" json:"role_id,omitempty"`
	Active      bool      `gorm:"default:true" json:"active"`
	System      bool      `gorm:"default:false" json:"system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
