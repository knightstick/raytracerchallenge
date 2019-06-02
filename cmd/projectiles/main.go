package main

import (
	"fmt"

	"github.com/knightstick/raytracerchallenge/tuples"
)

type projectile struct {
	Position tuples.Tuple
	velocity tuples.Tuple
}

type environment struct {
	gravity tuples.Tuple
	wind    tuples.Tuple
}

func main() {
	// Projectile starts one unit about the origin
	// Velocity is normalized to 1 unit/tick
	p := projectile{tuples.NewPoint(0, 1, 0), tuples.NewVector(0, 1, 0).Normalize()}

	// Gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e := environment{tuples.NewVector(0, -0.1, 0), tuples.NewVector(-0.01, 0, 0)}

	iteration := 0
	for {
		if p.Position.Y <= 0 {
			fmt.Printf("It took %d ticks for the projectile to hit the ground\n", iteration)
			break
		}

		fmt.Printf("Tick %d - Projectile is at %v\n", iteration, p.Position)

		p = tick(e, p)
		iteration++
	}
}

func tick(env environment, proj projectile) projectile {
	position := proj.Position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)

	return projectile{position, velocity}
}
