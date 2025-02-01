package setup

import (
	"fmt"

	"github.com/oadultradeepfield/gravigo/internal/simulator"
)

func InitializeSystem(cfg *InputConfig) ([]*simulator.Body, error) {
	var system []*simulator.Body

	for _, b := range cfg.Bodies {
		position, err := simulator.NewVector(b.Position[0], b.Position[1], b.Position[2], cfg.CoordinateType)
		if err != nil {
			return nil, fmt.Errorf("error reading position: %v", err)
		}

		velocity, err := simulator.NewVector(b.Velocity[0], b.Velocity[1], b.Velocity[2], cfg.CoordinateType)
		if err != nil {
			return nil, fmt.Errorf("error reading velocity: %v", err)
		}

		body, err := simulator.NewBody(b.Mass, position, velocity)
		if err != nil {
			return nil, fmt.Errorf("error creating body: %v", err)
		}

		system = append(system, body)
	}

	return system, nil
}
