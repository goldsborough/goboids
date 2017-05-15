package main

import (
	"flag"
	"fmt"
	"github.com/goldsborough/goboids/boids"
	"github.com/goldsborough/goboids/world"
	"os"
	"time"
)

var swarmSize = flag.Int("swarm-size", 20, "The size of the boid swarm")
var screenWidth = flag.Int("width", 80, "The number of columns to use for the screen")
var screenHeight = flag.Int("height", 20, "The number of lines to use for the screen")
var fps = flag.Int("fps", 10, "The frames to render per second")
var duration = flag.Int("duration", 10, "The number of seconds to render")

var cohesion = flag.Float64("cohesion", 0.5, "The cohesion constant")
var alignment = flag.Float64("alignment", 0.99, "The alignment constant")
var separation = flag.Float64("separation", 0.05, "The separation constant")
var bounce = flag.Float64("bounce", 4, "The bounce constant, for bouncing off walls")
var perturb = flag.Float64("perturb", 0.15, "The perturbation constant")
var center = flag.Float64("center", 0.1, "A constant defining how much to stay around the center")
var velocityClip = flag.Int("velocity-clip", 10, "The maximum velocity")


func main() {
	flag.Parse()

	world, err := world.New(*screenHeight, *screenWidth)
	if err != nil {
		fmt.Printf("Error setting up environment: %s", err)
		os.Exit(-1)
	}

	constants := boids.Constants{*cohesion, *alignment, *separation, *bounce, *center, *perturb, *velocityClip}
	swarm := boids.New(*swarmSize, world.SeedPoint(), constants)

	framePeriod := time.Duration(1000.0/float64(*fps)) * time.Millisecond
	numberOfFrames := *fps * *duration

	for i := 0; i < numberOfFrames; i++ {
		swarm.Update(world)
		world.Render(swarm)
		time.Sleep(framePeriod)
	}

}
