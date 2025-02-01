package simulator

import "fmt"

func HandleCollisions(bodies []*Body) error {
	for i := 0; i < len(bodies); i++ {
		b1 := bodies[i]

		for j := i + 1; j < len(bodies); j++ {
			b2 := bodies[j]

			_, _, _, distance, err := b1.Position.DistanceTo(b2.Position)
			if err != nil {
				return fmt.Errorf("error handling collision: %v", err)
			}

			if distance < CollisionThreshold {
				handleCollisionPair(b1, b2)
			}
		}
	}
	return nil
}

func handleCollisionPair(b1, b2 *Body) {
	b1.Velocity.E1, b2.Velocity.E1 = calculateElasticCollision(
		b1.Mass, b2.Mass, b1.Velocity.E1, b2.Velocity.E1,
	)

	b1.Velocity.E2, b2.Velocity.E2 = calculateElasticCollision(
		b1.Mass, b2.Mass, b1.Velocity.E2, b2.Velocity.E2,
	)

	b1.Velocity.E3, b2.Velocity.E3 = calculateElasticCollision(
		b1.Mass, b2.Mass, b1.Velocity.E3, b2.Velocity.E3,
	)
}

func calculateElasticCollision(m1, m2, v1, v2 float64) (newV1, newV2 float64) {
	totalMass := m1 + m2
	newV1 = ((m1-m2)*v1 + 2*m2*v2) / totalMass
	newV2 = ((m2-m1)*v2 + 2*m1*v1) / totalMass
	return
}
