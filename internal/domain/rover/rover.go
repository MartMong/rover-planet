package rover

import "fmt"

type Interface interface {
	TurnLeft() string
	TurnRight() string
	Forward() string
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

func New(mapSize int) Interface {
	return &Rover{
		mapSize:   mapSize,
		position:  &position{x: 0, y: 0},
		direction: N,
	}
}

func (r *Rover) TurnLeft() string {
	r.direction = (r.direction - 1 + 4) % 4
	return r.getState()
}

func (r *Rover) TurnRight() string {
	r.direction = (r.direction + 1 + 4) % 4
	return r.getState()
}

func (r *Rover) Forward() string {
	return r.step(1)
}

func (r *Rover) step(step int) (state string) {
	defer func() {
		state = r.getState()
	}()

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

	if x < 0 || x > r.mapSize {
		return
	}

	if y < 0 || y > r.mapSize {
		return
	}

	r.position.x = x
	r.position.y = y
	return
}

func (r *Rover) getState() string {
	fmt.Printf("%s:%d,%d\n", getDirection(r.direction), r.position.x, r.position.y)
	return fmt.Sprintf("%s:%d,%d", getDirection(r.direction), r.position.x, r.position.y)
}
