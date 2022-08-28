package rover

import "fmt"

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
	switch order {
	case "F":
		r.step(1)

	case "L":
		r.turnLeft()

	case "R":
		r.turnRight()

	default:
		return r.GetState(), fmt.Errorf("unsupport command: %s", order)
	}

	return r.GetState(), nil
}

func (r *Rover) GetState() string {
	return fmt.Sprintf("%s:%d,%d", getDirection(r.direction), r.position.x, r.position.y)
}

func (r *Rover) turnLeft() {
	r.direction = (r.direction - 1 + 4) % 4
}

func (r *Rover) turnRight() {
	r.direction = (r.direction + 1 + 4) % 4
}

func (r *Rover) step(step int) {
	x := r.position.x
	y := r.position.y

	switch r.direction {
	case N:
		y = y + step

	case E:
		x = x + step

	case S:
		y = y - step

	case W:
		x = x - step
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
