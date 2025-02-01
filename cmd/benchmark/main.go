package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/oadultradeepfield/gravigo/internal/simulator"
)

const (
	MinBodies             = 10
	MaxBodies             = 10000
	StepBodies            = 5
	GravitationalConstant = 1
	Dt                    = 0.01
	TotalTime             = 10.0
)

func main() {
	log.Println("Starting benchmark...")

	for n := MinBodies; n <= MaxBodies; n *= StepBodies {
		bodies, err := generateRandomBodies(n)
		if err != nil {
			log.Fatalf("Failed to generate bodies: %v", err)
		}

		log.Printf("Simulating %d bodies...", n)

		startTime := time.Now()
		simulator.RunSimulation(bodies, Dt, TotalTime, "benchmark.txt", GravitationalConstant)
		duration := time.Since(startTime)

		log.Printf("N=%d: Time taken: %s\n", n, duration.Round(time.Millisecond))
	}

	log.Println("Benchmark completed.")
}

func generateRandomBodies(n int) ([]*simulator.Body, error) {
	var system []*simulator.Body

	for i := 0; i < n; i++ {
		position, err := simulator.NewVector(rand.Float64(), rand.Float64(), rand.Float64(), simulator.Cartesian)
		if err != nil {
			return nil, fmt.Errorf("error creating position vector: %v", err)
		}

		velocity, err := simulator.NewVector(rand.Float64(), rand.Float64(), rand.Float64(), simulator.Cartesian)
		if err != nil {
			return nil, fmt.Errorf("error creating velocity vector: %v", err)
		}

		body, err := simulator.NewBody(0.1+rand.Float64()*(1.0-0.1), position, velocity)
		if err != nil {
			return nil, fmt.Errorf("error creating body: %v", err)
		}

		system = append(system, body)
	}
	return system, nil
}
