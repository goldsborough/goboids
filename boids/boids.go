package boids

import (
	"fmt"
	"github.com/goldsborough/goboids/point"
	"github.com/goldsborough/goboids/world"
	"math"
)

type Boid struct {
	id       int
	position point.Point
	velocity point.Point
}

type Constants struct {
  Cohesion, Alignment, Separation, Bounce, Center, Perturb float64
  VelocityClip int
}

type Swarm struct {
	boids []Boid
  constants Constants
  centerPoint point.Point
}


func New(size int, center point.Point, constants Constants) *Swarm {
	swarm := &Swarm{make([]Boid, size), constants, center}
	radius := int(math.Sqrt(float64(swarm.Size())))
	for index, _ := range swarm.boids {
		swarm.boids[index].id = index
		swarm.boids[index].position = point.RandomAround(center, radius)
	}

	return swarm
}

func (swarm *Swarm) Points() map[point.Point]bool {
	points := make(map[point.Point]bool, swarm.Size())
	for _, boid := range swarm.boids {
		points[boid.position] = true
	}
	return points
}

func (swarm *Swarm) Update(world *world.World) {
	swarm.group()
	swarm.separate()
	swarm.align()
  swarm.clip()
  swarm.center()
  swarm.bounce(world)
}

func (swarm *Swarm) group() {
	var center point.Point
	for _, boid := range swarm.boids {
		center = center.Add(boid.position)
	}

	for index, boid := range swarm.boids {
		perceivedCenter := center.Sub(boid.position).Div(swarm.Size() - 1)
		delta := perceivedCenter.Sub(boid.position)
		swarm.boids[index].velocity = boid.velocity.Add(delta.Scale(swarm.constants.Cohesion))
	}
}

func (swarm *Swarm) separate() {
	// TODO: Since boids are never added dynamically, we could sort the array and move only locally a
	// few positions left and right, quadratic for now

	for index, boid := range swarm.boids {
		for otherIndex, other := range swarm.boids {
			if index != otherIndex {
        swarm.boids[index].position.Perturb(swarm.constants.Perturb)
				distance := point.Distance(boid.position, other.position)
				if distance < 4 {
					delta := other.position.Sub(boid.position)
					swarm.boids[index].velocity = boid.velocity.Sub(delta.Scale(swarm.constants.Separation))
				}
			}
		}
	}
}

func (swarm *Swarm) align() {
	var center point.Point
	for _, boid := range swarm.boids {
		center = center.Add(boid.velocity)
	}

	for index, boid := range swarm.boids {
		perceivedDirection := center.Sub(boid.velocity).Div(swarm.Size() - 1)
		delta := perceivedDirection.Sub(boid.velocity)
		swarm.boids[index].velocity = boid.velocity.Add(delta.Scale(swarm.constants.Alignment))
	}
}

func (swarm *Swarm) bounce(world *world.World) {
  for index, boid := range swarm.boids {
    if ok := world.CheckBounds(boid.position.Add(boid.velocity)); !ok {
      swarm.boids[index].velocity = boid.velocity.Negate().Scale(swarm.constants.Bounce)
    }
    swarm.boids[index].Move()
  }
}

func (swarm *Swarm) clip() {
  for index, boid := range swarm.boids {
		if boid.velocity.Norm() > swarm.constants.VelocityClip {
      swarm.boids[index].velocity.X = swarm.constants.VelocityClip
      swarm.boids[index].velocity.Y = swarm.constants.VelocityClip
    }
	}
}

func (swarm *Swarm) center() {
  for index, boid := range swarm.boids {
		delta := swarm.centerPoint.Sub(boid.position)
		swarm.boids[index].velocity = boid.velocity.Add(delta.Scale(swarm.constants.Center))
	}
}

func (swarm *Swarm) Size() int {
	return len(swarm.boids)
}

func (boid Boid) String() string {
	return fmt.Sprintf("<Boid: %v | %v | %v>", boid.id, boid.position, boid.velocity)
}

func (boid *Boid) Points() []point.Point {
	return []point.Point{boid.position}
}

func (boid *Boid) Move() {
  boid.position.Move(boid.velocity)
}
