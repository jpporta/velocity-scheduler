package main

import (
	"fmt"
	"time"

	"github.com/jpporta/velocity-classes-scheduler/internal/integration/velocity"
	"github.com/jpporta/velocity-classes-scheduler/internal/models"
)

func main() {
	cfg := models.GetConfig()

	velocityInstance := velocity.Login(cfg.Users[0])
	// Fetch available classes
	fmt.Println(cfg.Unit)

	classes := velocityInstance.GetClasses(cfg.Unit.ID, time.Now())

	// Filter classes based on config and previous execution results

	// Possible: Resolve seats

	// Get into classes
}
