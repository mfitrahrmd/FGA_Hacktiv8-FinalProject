package domain

import "time"

type Base struct {
	Id        uint       `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
