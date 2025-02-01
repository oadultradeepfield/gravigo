package simulator

func RungeKuttaStep(bodies []*Body, dt, gravitationalConstant float64) {
	n := len(bodies)

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
			pos := b.Position
			vel := b.Velocity
			tmp[i], _ = NewBody(b.Mass, pos, vel)
		}
		return tmp
	}

	for i, b := range bodies {
		k1Pos[i] = b.Velocity
		b.UpdateAcceleration(bodies, gravitationalConstant)
		k1Vel[i] = b.Acceleration
	}

	tmp := copyBodies()
	for i := range tmp {
		tmp[i].Position.E1 += 0.5 * dt * k1Pos[i].E1
		tmp[i].Position.E2 += 0.5 * dt * k1Pos[i].E2
		tmp[i].Position.E3 += 0.5 * dt * k1Pos[i].E3
		tmp[i].Velocity.E1 += 0.5 * dt * k1Vel[i].E1
		tmp[i].Velocity.E2 += 0.5 * dt * k1Vel[i].E2
		tmp[i].Velocity.E3 += 0.5 * dt * k1Vel[i].E3
	}

	for i, b := range tmp {
		k2Pos[i] = b.Velocity
		b.UpdateAcceleration(tmp, gravitationalConstant)
		k2Vel[i] = b.Acceleration
	}

	tmp = copyBodies()
	for i := range tmp {
		tmp[i].Position.E1 += 0.5 * dt * k2Pos[i].E1
		tmp[i].Position.E2 += 0.5 * dt * k2Pos[i].E2
		tmp[i].Position.E3 += 0.5 * dt * k2Pos[i].E3
		tmp[i].Velocity.E1 += 0.5 * dt * k2Vel[i].E1
		tmp[i].Velocity.E2 += 0.5 * dt * k2Vel[i].E2
		tmp[i].Velocity.E3 += 0.5 * dt * k2Vel[i].E3
	}

	for i, b := range tmp {
		k3Pos[i] = b.Velocity
		b.UpdateAcceleration(tmp, gravitationalConstant)
		k3Vel[i] = b.Acceleration
	}

	tmp = copyBodies()
	for i := range tmp {
		tmp[i].Position.E1 += dt * k3Pos[i].E1
		tmp[i].Position.E2 += dt * k3Pos[i].E2
		tmp[i].Position.E3 += dt * k3Pos[i].E3
		tmp[i].Velocity.E1 += dt * k3Vel[i].E1
		tmp[i].Velocity.E2 += dt * k3Vel[i].E2
		tmp[i].Velocity.E3 += dt * k3Vel[i].E3
	}

	for i, b := range tmp {
		k4Pos[i] = b.Velocity
		b.UpdateAcceleration(tmp, gravitationalConstant)
		k4Vel[i] = b.Acceleration
	}

	for i, b := range bodies {
		b.Position.E1 += dt / 6.0 * (k1Pos[i].E1 + 2*k2Pos[i].E1 + 2*k3Pos[i].E1 + k4Pos[i].E1)
		b.Position.E2 += dt / 6.0 * (k1Pos[i].E2 + 2*k2Pos[i].E2 + 2*k3Pos[i].E2 + k4Pos[i].E2)
		b.Position.E3 += dt / 6.0 * (k1Pos[i].E3 + 2*k2Pos[i].E3 + 2*k3Pos[i].E3 + k4Pos[i].E3)
		b.Velocity.E1 += dt / 6.0 * (k1Vel[i].E1 + 2*k2Vel[i].E1 + 2*k3Vel[i].E1 + k4Vel[i].E1)
		b.Velocity.E2 += dt / 6.0 * (k1Vel[i].E2 + 2*k2Vel[i].E2 + 2*k3Vel[i].E2 + k4Vel[i].E2)
		b.Velocity.E3 += dt / 6.0 * (k1Vel[i].E3 + 2*k2Vel[i].E3 + 2*k3Vel[i].E3 + k4Vel[i].E3)
	}
}
