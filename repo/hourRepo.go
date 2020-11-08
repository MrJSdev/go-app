package repo

import (
	"MrHour/entity"
)

// HourRepo to retain functions
type HourRepo interface {
	FindAll() (entity.Hour, error)
	Save(hour entity.Hour) (entity.Hour, error)
}
