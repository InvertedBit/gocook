package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `gorm:"default:now()"`
	UpdatedAt time.Time      `gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type EntryType int

const (
	EntryTypeActive EntryType = iota
	EntryTypeEdit
	EntryTypeHistory
)

type CoWModel struct {
	Model
	EntryType  EntryType `gorm:"type:int;not null;default:0;index"`
	PreviousID *string   `gorm:"type:uuid;index;nullable"`
}
