package simulator

import (
	"fmt"
	"os"
)

func RunSimulation(bodies []*Body, dt, totalTime float64, filename string, gravitationalConstant float64) error {
	steps := int(totalTime / dt)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error clearing output file: %v", err)
	}
	file.Close()

	for i := 0; i < steps; i++ {
		RungeKuttaStep(bodies, dt, gravitationalConstant)
		HandleCollisions(bodies)
		printState(bodies, filename)
	}
	return nil
}

func printState(bodies []*Body, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filename, err)
	}
	defer file.Close()

	for _, body := range bodies {
		_, err := fmt.Fprintf(file, "%f, %f, %f\n", body.Position.E1, body.Position.E2, body.Position.E3)
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", filename, err)
		}
	}
	return nil
}
