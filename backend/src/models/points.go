package models

import "github.com/google/uuid"

type Points struct {
	BaseModel
	CompetitionID uuid.UUID `gorm:"type:uuid;column:competition_id"`
	UserID uuid.UUID `gorm:"type:uuid;column:user_id"`
}