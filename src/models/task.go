package models

import (
	"time"

	"github.com/AnD00/my-dear-tasks/src/enum"
)

type Task struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Text      string
	Status    enum.State
}
