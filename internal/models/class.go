package models

import (
	"github.com/jpporta/velocity-classes-scheduler/internal/utils"
)

type Class struct {
	Start string
	Duration int
	DayOfWeek utils.DayOfWeek
	Date string
	Instructor int
	OnlyForUsers []int
}
