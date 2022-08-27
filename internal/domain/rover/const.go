package rover

type direction int

const (
	N direction = iota
	E
	S
	W
)

func getDirection(d direction) string {
	switch d {
	case N:
		return "N"

	case E:
		return "E"

	case S:
		return "S"

	case W:
		return "W"

	default:
		return ""
	}
}
