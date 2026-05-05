package services

import (
	"time"

	"github.com/jpporta/velocity-classes-scheduler/internal/integration/velocity"
	"github.com/jpporta/velocity-classes-scheduler/internal/models"
	"github.com/jpporta/velocity-classes-scheduler/internal/utils"
)

func (c models.Class) FilterClasses(userID int, classes []velocity.VelocityClasses) []velocity.VelocityClasses {
	res := make([]velocity.VelocityClasses, 0)

	for _, class := range classes {
		if c.Match(userID, classes) {
		}
	}

	return res
}

func (c models.Class) Match(userID int, class velocity.VelocityClasses) bool {

	if c.Duration != 0 && c.Duration != class.Duration {
		return false
	}

	if c.DayOfWeek != utils.Unset && c.DayOfWeek != utils.DayOfWeek(class.Start.Weekday() + 1) {
		return false
	}

	if len(c.OnlyForUsers) > 0 && func() bool {
		for _, id := range c.OnlyForUsers {
			if id == userID {
				return false
			}
		}
		return true
	}() {
		return false
	}

	if c.Date != "" && c.Date != class.Start.Format(time.DateOnly) {
		return false
	}

	if c.Start != "" && c.Start != class.Start.Format(time.TimeOnly) {
		return false
	}

	if c.Instructor != 0 && c.Instructor != class.Instructor {
		return false
	}
	return true
}
