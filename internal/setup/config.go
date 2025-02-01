package setup

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oadultradeepfield/gravigo/internal/simulator"
)

type InputConfig struct {
	SimulatorConfig SimulatorConfig          `json:"simulator_config"`
	CoordinateType  simulator.CoordinateType `json:"coordinate_type"`
	Bodies          []BodyInput              `json:"bodies"`
}

type SimulatorConfig struct {
	GravitationalConstant float64 `json:"gravitational_constant"`
	Dt                    float64 `json:"dt"`
	TotalTime             float64 `json:"total_time"`
	OutputFile            string  `json:"output_file"`
}

type BodyInput struct {
	Name     string    `json:"_name"`
	Mass     float64   `json:"mass"`
	Position []float64 `json:"position"`
	Velocity []float64 `json:"velocity"`
}

func LoadConfig(filename string) (*InputConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error loading configuration file %s: %v", filename, err)
	}
	defer file.Close()

	var cfg InputConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("error decoding json configuration file: %v", err)
	}

	return &cfg, nil
}
