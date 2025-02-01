package simulator

import (
	"errors"
	"runtime"
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

	numWorkers := runtime.NumCPU()
	bodiesPerWorker := len(bodies) / numWorkers
	results := make([]Vector, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = len(bodies)
			}

			for j := start; j < end; j++ {
				other := bodies[j]
				if other == b {
					continue
				}
				dx, dy, dz, radialDistance, err := b.Position.DistanceTo(other.Position)
				if err != nil {
					continue
				}
				softenedDistanceSquared := radialDistance*radialDistance + GravitationalSoftening*GravitationalSoftening
				accelerationMagnitude := gravitationalConstant * other.Mass / softenedDistanceSquared
				results[workerID].E1 += accelerationMagnitude * dx / radialDistance
				results[workerID].E2 += accelerationMagnitude * dy / radialDistance
				results[workerID].E3 += accelerationMagnitude * dz / radialDistance
			}
		}(i)
	}

	wg.Wait()

	for _, result := range results {
		b.Acceleration.E1 += result.E1
		b.Acceleration.E2 += result.E2
		b.Acceleration.E3 += result.E3
	}

	return nil
}
