package models

import "github.com/google/uuid"

type Points struct {
	BaseModel
	CompetitionID uuid.UUID
	UserID uuid.UUID
}