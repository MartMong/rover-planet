package rover

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Interface interface {
	Command(order string) (string, error)
	GetState() string
}

type position struct {
	x int
	y int
}

type Rover struct {
	mapSize   int
	position  *position
	direction direction
}

func New(mapSize int) *Rover {
	return &Rover{
		mapSize:   mapSize,
		position:  &position{x: 0, y: 0},
		direction: N,
	}
}

func (r *Rover) Command(order string) (string, error) {
	switch {
	case order == "F":
		r.step(1)

	case regexp.MustCompile(`^F\d+`).MatchString(order):
		step, err := strconv.Atoi(order[1:])
		if err != nil {
			return r.GetState(), fmt.Errorf("unsupport command: %s", order)
		}
		r.step(step)

	case order == "B":
		r.step(-1)

	case regexp.MustCompile(`^B\d+`).MatchString(order):
		step, err := strconv.Atoi(order[1:])
		if err != nil {
			return r.GetState(), fmt.Errorf("unsupport command: %s", order)
		}
		r.step(-1 * step)

	case order == "L":
		r.turnLeft()

	case order == "R":
		r.turnRight()

	case order == "HL":
		r.turnHalfLeft()

	case order == "HR":
		r.turnHalfRight()

	default:
		return r.GetState(), fmt.Errorf("unsupport command: %s", order)
	}

	return r.GetState(), nil
}

func (r *Rover) GetState() string {
	return fmt.Sprintf("%s:%d,%d", getDirection(r.direction), r.position.x, r.position.y)
}

func (r *Rover) turnLeft() {
	r.direction = direction(math.Mod(float64(r.direction-1.0+4.0), 4.0))
}

func (r *Rover) turnRight() {
	r.direction = direction(math.Mod(float64(r.direction+1.0+4.0), 4.0))
}

func (r *Rover) turnHalfLeft() {
	r.direction = direction(math.Mod(float64(r.direction-0.5+4.0), 4.0))
}

func (r *Rover) turnHalfRight() {
	r.direction = direction(math.Mod(float64(r.direction+0.5+4.0), 4.0))
}

func (r *Rover) step(step int) {
	x := r.position.x
	y := r.position.y

	switch r.direction {
	case N:
		y = y + step

	case NE:
		x = x + step
		y = y + step

	case E:
		x = x + step

	case SE:
		x = x + step
		y = y - step

	case S:
		y = y - step

	case SW:
		x = x - step
		y = y - step

	case W:
		x = x - step

	case NW:
		x = x - step
		y = y + step
	}

	if x < 0 || x >= r.mapSize {
		return
	}

	if y < 0 || y >= r.mapSize {
		return
	}

	r.position.x = x
	r.position.y = y
	return
}
