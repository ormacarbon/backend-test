package models

import "github.com/google/uuid"

type BaseModel struct {
	ID        uuid.UUID `json:"id"        gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt int64     `json:"createdAt" gorm:"autoCreateTime"`
}
