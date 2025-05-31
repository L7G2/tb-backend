package models

import "time"

type WeightLog struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Weight    float64   `json:"weight" db:"weight"`
	BodyFat   *float64  `json:"body_fat,omitempty" db:"body_fat"`
	Notes     string    `json:"notes,omitempty" db:"notes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
