package models

import "github.com/google/uuid"

type Competitions struct {
	BaseModel
	Status bool `json:"status" gorm:"type:boolean;not null;column:status"`
}

func (comp *Competitions) IsEmpty() bool {
	return comp.ID == uuid.Nil
}
