package main

import (
	"fmt"

	"github.com/knightstick/raytracerchallenge/tuples"
)

type projectile struct {
	Position tuples.Point
	velocity tuples.Vector
}

type environment struct {
	gravity tuples.Vector
	wind    tuples.Vector
}

func main() {
	// Projectile starts one unit about the origin
	// Velocity is normalized to 1 unit/tick
	position := tuples.NewPoint(0, 1, 0)
	velocity := tuples.Vector(tuples.Normalize(tuples.New(0, 1, 0, 0)).(tuples.Tuple))

	p := projectile{position, velocity}

	// Gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e := environment{tuples.NewVector(0, -0.1, 0), tuples.NewVector(-0.01, 0, 0)}

	iteration := 0
	for {
		if p.Position.Y() <= 0 {
			fmt.Printf("It took %d ticks for the projectile to hit the ground\n", iteration)
			break
		}

		fmt.Printf("Tick %d - Projectile is at %v\n", iteration, p.Position)

		p = tick(e, p)
		iteration++
	}
}

func tick(env environment, proj projectile) projectile {
	position := tuples.Point(tuples.Add(proj.Position, proj.velocity).(tuples.Tuple))
	velocity := tuples.Vector(tuples.Add(env.wind, tuples.Add(proj.velocity, env.gravity)).(tuples.Tuple))

	return projectile{position, velocity}
}
