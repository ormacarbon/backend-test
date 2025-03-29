package models

import "github.com/google/uuid"

type Competitions struct {
	BaseModel
	Status bool `json:"status" gorm:"type:boolean;not null;column:status"`
}

type CompetitionReport struct {
	Name string `json:"name" gorm:"column:name"`
	Points uint64 `json:"points" gorm:"column:points"`
	Email string `json:"-" gorm:"column:email"` 
	Phone string `json:"-" gorm:"column:phone"`
}

func (comp *Competitions) IsEmpty() bool {
	return comp.ID == uuid.Nil
}
