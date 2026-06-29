package document

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)
/*
GORM
Models
AutoMigrate
CRUD
Transactions
before create hook in orm
*/

type Document struct {
	//gorm.Model

	Id uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Version int `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	if d.Id == uuid.Nil {
		d.Id = uuid.New()
	}
	return nil
}




