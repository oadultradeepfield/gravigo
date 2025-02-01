package simulator

import (
	"errors"
	"log"
	"sync"
)

const (
	CollisionThreshold     = 1e-8
	GravitationalSoftening = 1e-8
)

type Body struct {
	Mass         float64
	Position     *Vector
	Velocity     *Vector
	Acceleration *Vector
}

func NewBody(mass float64, position, velocity *Vector) (*Body, error) {
	if mass <= 0 {
		return nil, errors.New("mass must be a positive value")
	}
	if position.Type != Cartesian || velocity.Type != Cartesian {
		ConvertSphericalToCartesian(position, velocity)
	}
	return &Body{
		Mass:         mass,
		Position:     position,
		Velocity:     velocity,
		Acceleration: &Vector{Type: Cartesian},
	}, nil
}

func (b *Body) UpdateAcceleration(bodies []*Body, gravitationalConstant float64) error {
	if gravitationalConstant <= 0 {
		return errors.New("gravitational constant must be a positive value")
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, other := range bodies {
		if other == b {
			continue
		}

		wg.Add(1)
		go func(other *Body) {
			defer wg.Done()

			dx, dy, dz, radialDistance, err := b.Position.DistanceTo(other.Position)
			if err != nil {
				log.Fatalf("error updating acceleration: %v", err)
				return
			}

			softenedDistanceSquared := radialDistance*radialDistance + GravitationalSoftening*GravitationalSoftening
			accelerationMagnitude := gravitationalConstant * other.Mass / softenedDistanceSquared

			mu.Lock()
			b.Acceleration.E1 += accelerationMagnitude * dx / radialDistance
			b.Acceleration.E2 += accelerationMagnitude * dy / radialDistance
			b.Acceleration.E3 += accelerationMagnitude * dz / radialDistance
			mu.Unlock()
		}(other)
	}

	wg.Wait()
	return nil
}
