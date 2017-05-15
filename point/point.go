package point

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (point Point) String() string {
	return fmt.Sprintf("(x = %d, y = %d)", point.X, point.Y)
}

func (point Point) Add(other Point) Point {
	point.X += other.X
	point.Y += other.Y
	return point
}

func (point *Point) Move(other Point) {
  *point = point.Add(other)
}

func (point Point) Sub(other Point) Point {
	point.X -= other.X
	point.Y -= other.Y
	return point
}

func (point Point) Div(scalar int) Point {
	point.X /= scalar
	point.Y /= scalar
	return point
}

func (point Point) Scale(scalar float64) Point {
	point.X = int(float64(point.X) * scalar)
	point.Y = int(float64(point.Y) * scalar)
	return point
}

func (point Point) Negate() Point {
  point.X = -point.X
  point.Y = -point.Y
  return point
}

func RandomAround(center Point, radius int) Point {
	x := center.X + (rand.Intn(2*radius) - radius)
	y := center.Y + (rand.Intn(2*radius) - radius)
	return Point{x, y}
}

func (point Point) Norm() int {
  dx := math.Pow(float64(point.X), 2)
  dy := math.Pow(float64(point.Y), 2)
	return int(math.Sqrt(dx + dy))
}

func Distance(first, second Point) int {
  dx := math.Pow(float64(first.X - second.X), 2)
  dy := math.Pow(float64(first.Y - second.Y), 2)
	return int(math.Sqrt(dx + dy))
}

func round(value float64) int {
  if math.Abs(value) < 0.5 {
    return 0
  } else if value < 0 {
    return -1
  } else {
    return +1
  }
}

func (point *Point) Perturb(constant float64) {
  point.X += round(rand.NormFloat64() * constant)
	point.Y += round(rand.NormFloat64() * constant)
}

// pixel = points, color, character
