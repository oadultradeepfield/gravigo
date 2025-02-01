package simulator

import (
	"runtime"
	"sync"
)

type CollisionPair struct {
	b1, b2 *Body
}

func HandleCollisions(bodies []*Body) error {
	numWorkers := runtime.NumCPU()
	totalPairs := (len(bodies) * (len(bodies) - 1)) / 2
	pairsPerWorker := totalPairs / numWorkers
	var wg sync.WaitGroup
	CollisionPairs := make([][]CollisionPair, numWorkers)

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			pairStart := workerID * pairsPerWorker
			pairEnd := pairStart + pairsPerWorker
			if workerID == numWorkers-1 {
				pairEnd = totalPairs
			}

			pairCount := 0
			for i := 0; i < len(bodies); i++ {
				for j := i + 1; j < len(bodies); j++ {
					if pairCount < pairStart {
						pairCount++
						continue
					}
					if pairCount >= pairEnd {
						break
					}

					_, _, _, distance, err := bodies[i].Position.DistanceTo(bodies[j].Position)
					if err != nil {
						continue
					}

					if distance < CollisionThreshold {
						CollisionPairs[workerID] = append(CollisionPairs[workerID],
							CollisionPair{b1: bodies[i], b2: bodies[j]})
					}
					pairCount++
				}
			}
		}(w)
	}

	wg.Wait()

	for _, workerPairs := range CollisionPairs {
		for _, pair := range workerPairs {
			handleCollisionPair(pair.b1, pair.b2)
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
