package entities

import (
	"gorm.io/gorm"
	"time"
)

type Content struct {
	gorm.Model
	Created time.Time
	Name    string
	Content []byte
}
