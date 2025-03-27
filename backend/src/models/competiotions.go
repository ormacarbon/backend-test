package models

type Competitions struct {
	BaseModel
	Status bool `json:"status" gorm:"type:boolean;not null;column:status"`
}
