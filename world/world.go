package world

import (
	"fmt"
	"github.com/goldsborough/goboids/point"
	"os"
	"os/exec"
)

type Renderable interface {
  Points() map[point.Point]bool
}

type World struct {
	height, width int
}

func New(maximumLines, maximumColumns int) (*World, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var lines, columns int
	_, err = fmt.Sscan(string(out), &lines, &columns)
	if err != nil {
		return nil, err
	}

  if lines > maximumLines {
    lines = maximumLines
  }
  if columns > maximumColumns {
    columns = maximumColumns
  }

	return &World{lines, columns}, nil
}

func (world *World) Render(object Renderable) {
  points := object.Points()
  for y := 0; y < world.height; y++ {
    for x := 0; x < world.width; x++ {
      if _, ok := points[point.Point{x, y}]; ok {
        fmt.Print("\033[91m🍟\033[0m");
      } else if (x == 0 || x == world.width - 1) {
        // fmt.Print("|");
      } else if y == 0 || y == world.height - 1 {
        // fmt.Print("-");
      } else {
        fmt.Print(" ");
      }
    }
    fmt.Println();
  }
  fmt.Printf("\r\033[%dA", world.height)
}

func (world *World) SeedPoint() point.Point {
  return point.Point{world.width/2, world.height/2}
}

func (world *World) CheckBounds(point point.Point) bool {
  if (point.X < 0 || point.X >= world.width) {
    return false
  }
  if (point.Y < 0 || point.Y >= world.height) {
    return false
  }
  return true
}

func (world *World) String() string {
  return fmt.Sprintf("<World: height = %d, width = %d>", world.height, world.width);
}
