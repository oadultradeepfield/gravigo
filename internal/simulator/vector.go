package simulator

import (
	"errors"
	"fmt"
	"math"
)

type CoordinateType string

const (
	Cartesian CoordinateType = "cartesian"
	Spherical CoordinateType = "spherical"
)

type Vector struct {
	E1, E2, E3 float64
	Type       CoordinateType
}

func NewVector(e1, e2, e3 float64, coordinateType CoordinateType) (*Vector, error) {
	if coordinateType != Cartesian && coordinateType != Spherical {
		return nil, fmt.Errorf("invalid coordinate type: %v", coordinateType)
	}
	return &Vector{E1: e1, E2: e2, E3: e3, Type: coordinateType}, nil
}

func (v *Vector) DistanceTo(other *Vector) (dx, dy, dz, distance float64, err error) {
	if v.Type != Cartesian || other.Type != Cartesian {
		return 0, 0, 0, 0, errors.New("distance calculation is only supported for Cartesian coordinates")
	}
	dx = other.E1 - v.E1
	dy = other.E2 - v.E2
	dz = other.E3 - v.E3
	distance = math.Sqrt(dx*dx + dy*dy + dz*dz)
	err = nil
	return
}

func (v *Vector) DeepCopy() *Vector {
	return &Vector{v.E1, v.E2, v.E3, Cartesian}
}

func ConvertSphericalToCartesian(position, velocity *Vector) error {
	if position.Type != Spherical || velocity.Type != Spherical {
		return errors.New("input vectors must be of Spherical coordinate type")
	}

	r := position.E1
	theta := position.E2
	phi := position.E3

	vr := velocity.E1
	vtheta := velocity.E2
	vphi := velocity.E3

	position.E1 = r * math.Cos(theta) * math.Sin(phi)
	position.E2 = r * math.Sin(theta) * math.Sin(phi)
	position.E3 = r * math.Cos(phi)
	position.Type = Cartesian

	velocity.E1 = vr*math.Sin(phi)*math.Cos(theta) - vtheta*math.Sin(theta) + vphi*math.Cos(theta)*math.Cos(phi)
	velocity.E2 = vr*math.Sin(phi)*math.Sin(theta) + vtheta*math.Cos(theta) + vphi*math.Cos(theta)*math.Sin(phi)
	velocity.E3 = vr*math.Cos(phi) - vphi*math.Sin(phi)
	velocity.Type = Cartesian

	return nil
}
