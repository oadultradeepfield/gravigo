package main

import (
	"log"
	"os"
	"time"

	"github.com/oadultradeepfield/gravigo/internal/setup"
	"github.com/oadultradeepfield/gravigo/internal/simulator"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <config_file.json>")
	}

	configFile := os.Args[1]
	cfg, err := setup.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	bodySystem, err := setup.InitializeSystem(cfg)
	if err != nil {
		log.Fatalf("error initializing system: %v", err)
	}

	log.Printf("Simulation started with %d bodies...\n", len(bodySystem))

	startTime := time.Now()

	simulator.RunSimulation(
		bodySystem,
		cfg.SimulatorConfig.Dt,
		cfg.SimulatorConfig.TotalTime,
		cfg.SimulatorConfig.OutputFile,
		cfg.SimulatorConfig.GravitationalConstant,
	)

	duration := time.Since(startTime)

	log.Printf("Simulation completed. Results saved to %s\n", cfg.SimulatorConfig.OutputFile)
	log.Printf("Time taken: %s\n", duration.Round(time.Millisecond))
}
