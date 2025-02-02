package simulator

import (
	"runtime"
	"sync"
)

func RungeKuttaStep(bodies []*Body, dt, gravitationalConstant float64) {
	n := len(bodies)
	numWorkers := runtime.NumCPU()
	bodiesPerWorker := n / numWorkers
	if bodiesPerWorker < 1 {
		numWorkers = n
		bodiesPerWorker = 1
	}

	k1Pos := make([]*Vector, n)
	k1Vel := make([]*Vector, n)
	k2Pos := make([]*Vector, n)
	k2Vel := make([]*Vector, n)
	k3Pos := make([]*Vector, n)
	k3Vel := make([]*Vector, n)
	k4Pos := make([]*Vector, n)
	k4Vel := make([]*Vector, n)

	copyBodies := func() []*Body {
		tmp := make([]*Body, n)
		for i, b := range bodies {
			tmp[i] = b.DeepCopy()
		}
		return tmp
	}

	processChunk := func(start, end int, bodies []*Body, fn func(int, *Body)) {
		for i := start; i < end; i++ {
			fn(i, bodies[i])
		}
	}

	var wg sync.WaitGroup

	// K1 calculation
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, bodies, func(i int, b *Body) {
				k1Pos[i] = b.Velocity.DeepCopy()
				b.UpdateAcceleration(bodies, gravitationalConstant)
				k1Vel[i] = b.Acceleration.DeepCopy()
			})
		}(w)
	}
	wg.Wait()

	// K2 preparation
	tmp := copyBodies()
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				b.Position.E1 += 0.5 * dt * k1Pos[i].E1
				b.Position.E2 += 0.5 * dt * k1Pos[i].E2
				b.Position.E3 += 0.5 * dt * k1Pos[i].E3
				b.Velocity.E1 += 0.5 * dt * k1Vel[i].E1
				b.Velocity.E2 += 0.5 * dt * k1Vel[i].E2
				b.Velocity.E3 += 0.5 * dt * k1Vel[i].E3
			})
		}(w)
	}
	wg.Wait()

	// K2 calculation
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				k2Pos[i] = b.Velocity.DeepCopy()
				b.UpdateAcceleration(tmp, gravitationalConstant)
				k2Vel[i] = b.Acceleration.DeepCopy()
			})
		}(w)
	}
	wg.Wait()

	// K3 preparation
	tmp = copyBodies()
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				b.Position.E1 += 0.5 * dt * k2Pos[i].E1
				b.Position.E2 += 0.5 * dt * k2Pos[i].E2
				b.Position.E3 += 0.5 * dt * k2Pos[i].E3
				b.Velocity.E1 += 0.5 * dt * k2Vel[i].E1
				b.Velocity.E2 += 0.5 * dt * k2Vel[i].E2
				b.Velocity.E3 += 0.5 * dt * k2Vel[i].E3
			})
		}(w)
	}
	wg.Wait()

	// K3 calculation
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				k3Pos[i] = b.Velocity.DeepCopy()
				b.UpdateAcceleration(tmp, gravitationalConstant)
				k3Vel[i] = b.Acceleration.DeepCopy()
			})
		}(w)
	}
	wg.Wait()

	// K4 preparation
	tmp = copyBodies()
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				b.Position.E1 += dt * k3Pos[i].E1
				b.Position.E2 += dt * k3Pos[i].E2
				b.Position.E3 += dt * k3Pos[i].E3
				b.Velocity.E1 += dt * k3Vel[i].E1
				b.Velocity.E2 += dt * k3Vel[i].E2
				b.Velocity.E3 += dt * k3Vel[i].E3
			})
		}(w)
	}
	wg.Wait()

	// K4 calculation
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, tmp, func(i int, b *Body) {
				k4Pos[i] = b.Velocity.DeepCopy()
				b.UpdateAcceleration(tmp, gravitationalConstant)
				k4Vel[i] = b.Acceleration.DeepCopy()
			})
		}(w)
	}
	wg.Wait()

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			start := workerID * bodiesPerWorker
			end := start + bodiesPerWorker
			if workerID == numWorkers-1 {
				end = n
			}
			processChunk(start, end, bodies, func(i int, b *Body) {
				b.Position.E1 += dt / 6.0 * (k1Pos[i].E1 + 2*k2Pos[i].E1 + 2*k3Pos[i].E1 + k4Pos[i].E1)
				b.Position.E2 += dt / 6.0 * (k1Pos[i].E2 + 2*k2Pos[i].E2 + 2*k3Pos[i].E2 + k4Pos[i].E2)
				b.Position.E3 += dt / 6.0 * (k1Pos[i].E3 + 2*k2Pos[i].E3 + 2*k3Pos[i].E3 + k4Pos[i].E3)
				b.Velocity.E1 += dt / 6.0 * (k1Vel[i].E1 + 2*k2Vel[i].E1 + 2*k3Vel[i].E1 + k4Vel[i].E1)
				b.Velocity.E2 += dt / 6.0 * (k1Vel[i].E2 + 2*k2Vel[i].E2 + 2*k3Vel[i].E2 + k4Vel[i].E2)
				b.Velocity.E3 += dt / 6.0 * (k1Vel[i].E3 + 2*k2Vel[i].E3 + 2*k3Vel[i].E3 + k4Vel[i].E3)
			})
		}(w)
	}
	wg.Wait()

	HandleCollisions(bodies)
}
